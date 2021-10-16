package file

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"path"
	"runtime/debug"
	"strings"
	"sync"

	"awesomeProject2/command"
	"awesomeProject2/handler"
)

const (
	FileError = "error in opening file"
	ExitFlag  = "exit"
	ReadFlag  = "done"
)

var (
	fileInstance file
	once         sync.Once
)

type file struct {
	readLocation  *os.File
	writeLocation *os.File
	readWriter    *bufio.ReadWriter
	delim         string
	handler       handler.HandleProvider
}

func (f file) Read() string {
	line, err := f.readWriter.ReadString([]byte(f.delim)[0])
	if err == io.EOF {
		return ExitFlag
	}

	// to add command in output file
	f.readWriter.WriteString(line)
	f.readWriter.Flush()

	line = strings.Replace(line, f.delim, "", -1)
	switch line {
	case "exit", "":
		return ExitFlag
	default:
		cmd := strings.Split(line, " ")
		systemStdOutFile := os.Stdout
		os.Stdout = f.writeLocation
		if err := f.handler.Get(cmd[0]).RunHandle(cmd[1:]...); err != nil {
			f.Write(err.Error())
		}
		os.Stdout = systemStdOutFile
	}
	return ReadFlag
}

func (f file) Write(s string) {
	_, err := f.readWriter.WriteString(s)
	if err != nil {
		log.Fatal(err.Error())
	}
	f.readWriter.WriteRune('\n')
	f.readWriter.Flush()
}

func (f file) Run() {
	defer func() {
		f.readLocation.Close()
		f.writeLocation.Close()
		if r := recover(); r != nil {
			debug.PrintStack()
		}
	}()
	for {
		if f.Read() == ExitFlag {
			f.Write("bye :( ")
			break
		}
	}
}

func NewClient(handleProvider handler.HandleProvider, params ...string) (command.Handler, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			if fileInstance.readLocation != nil {
				fileInstance.readLocation.Close()
			}
			if fileInstance.writeLocation != nil {
				fileInstance.writeLocation.Close()
			}
			debug.PrintStack()
		}
	}()
	once.Do(func() {
		fileInstance.handler = handleProvider
		if len(params) > 0 {
			fileInstance.readLocation, _ = os.OpenFile(params[0], os.O_RDONLY, 777)
			if err != nil {
				log.Fatal(err)
			}
			fileInstance.writeLocation, _ = os.Create(path.Join(path.Dir(params[0]), strings.TrimSuffix(path.Base(params[0]), path.Ext(params[0]))+"_output"+path.Ext(params[0])))
			if err != nil {
				log.Fatal(err)
			}
			fileInstance.readWriter = bufio.NewReadWriter(bufio.NewReader(fileInstance.readLocation), bufio.NewWriter(fileInstance.writeLocation))
		}
		if len(params) > 1 {
			fileInstance.delim = params[1]
		}
	})
	if err != nil {
		return nil, errors.New(FileError)
	}
	return fileInstance, nil
}