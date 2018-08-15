# todo-grpc
a sample grpc application

## Installation
*Prerequisites:*
- [go@v1.10](https://golang.org/doc/install)
- [dep@v0.4](https://github.com/golang/dep)
- [retool@v0.82](https://github.com/twitchtv/retool)
- [protoc@v3](https://github.com/google/protobuf)
- [docker@1.17](https://store.docker.com/search?type=edition&offering=community)
- [a configured go workspace](https://golang.org/doc/code.html)

```shell
# clone this repository
$ git clone https://github.com/Mindflash/todo-grpc $GOPATH/src/todo-grpc

# install dependencies and sync vendored executables
$ cd $GOPATH/src/todo-grpc && dep ensure -vendor-only && retool sync
```

## Getting Started
1. [Build the application](#Building)
2. Run the application with docker
    ```shell
    $ docker-compose up -d elasticsearch && docker-compose up --scale todo-grpc=3
    ```

## Documentation
- View the source code documentation at [localhost:6060/pkg/todo-grpc/](http://localhost:6060/pkg/todo-grpc/) with `godoc`:
    ```shell
    $ godoc -http=:6060
    ```

- View the protobuf source at [proto/todo.proto](proto/todo.proto):


- View the grpc api documentation in your default browser:
    ```shell
    # on macos
    $ open ./docs/index.html

    # on linux
    $ xdg-open ./docs/index.html
    ```

- View the json-over-http api documentation at [localhost:8080](http://localhost:8080) with `docker`:
    ```shell
    $ docker run -p 8080:8080 -e SWAGGER_JSON=/todo/todo.swagger.json -v $GOPATH/src/todo-grpc/docs:/todo -it swaggerapi/swagger-ui
    ```

## Code Generation
```shell
# set protoc path pointing to the root protoc directory (containing bin, include)
$ export PROTOC_PATH=/path/to/protoc

# run code gen via protoc
$ retool do protoc \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I $GOPATH/src \
    -I $PROTOC_PATH \
    -I $GOPATH/src/todo-grpc/proto \
    --go_out=plugins=grpc:$GOPATH/src/todo-grpc/proto \
    --grpc-gateway_out=logtostderr=true:$GOPATH/src/todo-grpc/proto \
    --swagger_out=logtostderr=true:$GOPATH/src/todo-grpc/docs \
    --doc_out=$GOPATH/src/todo-grpc/docs \
    --doc_opt=html,index.html \
    todo.proto

# do post processing
$ retool do protoc-go-inject-tag -input=./proto/todo.pb.go
```

## Building
```shell
$ retool do goreleaser --snapshot --rm-dist --skip-publish
```
## License
**UNLICENSED**

Copyright (c) 2018 Chris Ludden