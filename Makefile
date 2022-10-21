all: build
	@echo "Build complete! Executable is in bin/ directory."
	
test:
	go test -v ./test

build:
	mkdir -p bin
	go build -o bin/$(NAME) main.go

clean:
	rm -rf bin

.PHONY: clean

