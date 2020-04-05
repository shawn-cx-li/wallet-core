build-cli:
	go build cli/*.go

test:
	go test ./... -coverprofile coverage.out
	go tool cover -html=coverage.out
