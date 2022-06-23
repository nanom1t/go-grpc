# go-grpc

1. Install gRPC compiler:
```
sudo apt install protobuf-compiler
sudo apt install golang-goprotobuf-dev
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

2. Compile protobuf files:
```
protoc -I location/ location/location.proto --go_out=plugins=grpc:location
```

3. Run gRPC server:
```
$ go run server.go
2022/06/23 14:39:20 Starting gRPC server...
2022/06/23 14:39:36 Received distance GRPC request:
2022/06/23 14:39:36 -- Point 1: latitude:50.450001 longitude:30.523333 
2022/06/23 14:39:36 -- Point 2: latitude:49.842957 longitude:24.031111 
2022/06/23 14:52:13 Received distance GRPC request:
2022/06/23 14:52:13 -- Point 1: latitude:48.5522 longitude:24.4238 
2022/06/23 14:52:13 -- Point 2: latitude:49.842957 longitude:24.031111 
2022/06/23 14:52:13 Received distance GRPC request:
2022/06/23 14:52:13 -- Point 1: latitude:48.5522 longitude:24.4238 
2022/06/23 14:52:13 -- Point 2: latitude:50.450001 longitude:30.523333 
2022/06/23 14:53:55 Received distance GRPC request:
2022/06/23 14:53:55 -- Point 1: latitude:48.5522 longitude:24.4238 
2022/06/23 14:53:55 -- Point 2: latitude:49.842957 longitude:24.031111 
2022/06/23 14:53:55 Received distance GRPC request:
2022/06/23 14:53:55 -- Point 1: latitude:48.5522 longitude:24.4238 
2022/06/23 14:53:55 -- Point 2: latitude:50.450001 longitude:30.523333
```

4. Run gRPC client:
```
$ go run client.go
Distance to Lviv: 146333.745285
Distance to Kiev: 488209.947259
```
