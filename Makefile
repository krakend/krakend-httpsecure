all: test

test:
	go get .
	go test -cover ./...