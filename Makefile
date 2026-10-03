.PHONY: tidy fmt lint test bench

tidy:
	go mod tidy

fmt:
	gofmt -w .

lint:
	golangci-lint run ./...

test:
	go test -v -race ./...

bench:
	go test -bench=. -benchmem ./...
