#!/bin/bash
protoc  --go_out=./  \
        --go-grpc_out=./ \
        optionhub.v1.proto

protoc  --go_out=./  \
        --go-grpc_out=./ \
        new_attribute.proto

protoc --doc_out=. --doc_opt=markdown,README.md ./optionhub.proto
protoc --doc_out=. --doc_opt=markdown,README_v1.md ./optionhub.v1.proto ./new_attribute.proto