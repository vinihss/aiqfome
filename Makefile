GOCMD = go

build:
	go build -o bin/main cmd/server/main.go

run:
	go run cmd/server/main.go

docs:
	swag init --dir=cmd/server,internal

# Run linters and checks
vet:
	$(GOCMD) vet ./...


# Format code
fmt:
	$(GOCMD) fmt ./...


