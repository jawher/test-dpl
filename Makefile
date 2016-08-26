default: release

deps:
	go get -t -v .

release: deps
	GOOS=linux  GOARCH=amd64 go build -o bin/linux/test
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/test

.PHONY: deps release