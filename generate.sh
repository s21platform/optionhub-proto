#!/bin/bash
protoc  --go_out=./  \
        --go-grpc_out=./ \
        optionhub.proto

protoc --doc_out=. --doc_opt=markdown,README.md ./optionhub.proto