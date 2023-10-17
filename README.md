# decipher
  Scratch repo setup with file/cli communication capability.

### structure
  * handler: have implementation in handler.go and handlers in handlers.go
  * command: have definition in command.go
      * cli/cli.go: cli command implementation
      * file/file.go: file command implementation

## Usage
  1. Create handler instance using `GetHandles` that returns HandleProvider.
  2. Create command implementor client(cli/file client, say _commander_), using HandleProvider and delimiter config(Ex: "\n").
  3. Inject command implementor client(say _commander_) to command client that returns command handler.
  4. Use command handler `Run` method to run the commander.
