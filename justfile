# https://just.systems

default:
		go run cmd/carnet/main.go

test:
    go test -v ./internal/core
