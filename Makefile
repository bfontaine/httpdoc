SRCS=main.go $(wildcard httpdoc/*.go)

bin/httpdoc: $(SRCS)
	go build -o $@ .

check:
	go test -v ./...
