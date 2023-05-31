build:
	@go build -o bin/gof
run: build
	@./bin/gof
test:
	@go test -v ./..