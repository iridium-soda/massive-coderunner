all: build

build:
	go build github.com/iridium-soda/massive-coderunner/cmd/astool

install:
	go install github.com/iridium-soda/massive-coderunner/cmd/astool
clean:
	go clean
	rm -fr gen
