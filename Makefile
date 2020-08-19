.DEFAULT_GOAL := default

default: all

bin/protoform:
	go build -o bin/protoform cmd/main.go

clean:
	rm -f bin/*

all: clean bin/protoform


.PHONY: clean all
