generate-mocks:
	go generate ./...

test: 
	go test ./... -cover

test-debug: 
	go test ./... -cover -v