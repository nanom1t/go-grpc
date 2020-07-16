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
2020/07/16 14:36:08 Starting gRPC server...
2020/07/16 14:36:27 Received distance GRPC request:
2020/07/16 14:36:27 -- Point 1: latitude:50.450001 longitude:30.523333
2020/07/16 14:36:27 -- Point 2: latitude:49.842957 longitude:24.031111
```

4. Run gRPC client:
```
$ go run client.go
GRPC response: 467357.25793234253
```
