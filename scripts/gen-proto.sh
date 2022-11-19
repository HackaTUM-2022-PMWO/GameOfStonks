# !/bin/bash

# generate frontend protos
protoc --plugin=./frontend/node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=./frontend/src/proto ./proto/messages.proto

# generate backend protos
protoc --go_out=./backend/proto --go_opt=paths=source_relative \
    --go-grpc_out=./backend/proto --go-grpc_opt=paths=source_relative \
    proto/messages.proto
