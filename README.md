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
