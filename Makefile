build:
	go build -o limgo cmd/main.go

test:	build
	go test ./... -coverprofile=cov.out -race
	./limgo -covfile=cov.out -v=4
	rm limgo cov.out
	