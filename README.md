# todo-grpc
a sample grpc application

## Getting Started
```shell
$ protoc \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I $GOPATH/src -I /Users/cludden/protoc/ \
    -I $GOPATH/todo-grpc/proto \
    --go_out=plugins=grpc:. \
    --grpc-gateway_out=logtostderr=true:. \
    todo.proto 
```