fmt:
	gofmt -w .

lint:
	golangci-lint run

build:	fmt lint
	go build -o limgo cmd/main.go

test: 
	go test ./... -coverprofile=test.cov -race

test-cov: test
	go run cmd/main.go -covfile=test.cov -v=2
