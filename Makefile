#!/bin/bash

.PHONY:
gen: clean
	buf generate --path proto/helloworld

.PHONY:
fmt:
	go fmt ./...

.PHONY:
install:
	go get \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/bufbuild/buf/cmd/buf

.PHONY:
run:
	go run server/server.go -port 8080

.PHONY:
clean:
	rm -rf gen/go/*/*.go
