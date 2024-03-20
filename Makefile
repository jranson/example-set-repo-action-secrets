.PHONY: build clean run 

clean:
	@rm -rf bin

build: clean
	@go build -o bin/set-repo-secrets cmd/cli/main.go
