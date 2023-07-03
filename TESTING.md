## Clearing the Cache
go clean -cache

## Profiling the test coverage
go test -coverprofile ./tests/scheduler.out ./...
go tool cover -html=./tests/scheduler.out -o ./tests/scheduler.html
