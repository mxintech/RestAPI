# GOCC is a go compiler as variable
GOCC = go build

default: Build.app

Build.app: main.go
	$(GOCC) main.go

# Start the application
run:
	./main

# Go run command
gr:
	go run main.go

# Clean removes al binaries from app
clean:
	rm -rf *./main
