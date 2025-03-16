#!/bin/bash
protoc  --go_out=./  \
        --go-grpc_out=./ \
        optionhub.v8.proto

protoc --doc_out=. --doc_opt=markdown,README.md ./optionhub.proto ./optionhub.v8.proto