# Commands

### Initialise a go module
go mod init

### Update go module to import dependency from local system (non published dependency)
go mod edit -replace example.com/greetings=../greetings

go mod tidy

go run .

