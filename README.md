# wps store micro service

# protoc-gen
mv protoc-gen-swagger /usr/local/go/bin
mv protoc-gen-grpc-gateway /usr/local/go/bin

# 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

#编译rpc.proto为rpc.pb.proto
protoc -I/usr/local/include -I. \-I$GOPATH/src \-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \--go_out=plugins=grpc,Mgoogle/api/annotations.proto=wps_store/rpc/google/api:. \rpc.proto

#编译rpc.proto为rpc.pb.gw.proto
protoc -I/usr/local/include -I. \-I$GOPATH/src \-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \--grpc-gateway_out=logtostderr=true:. \rpc.proto

#编译 swagger.json
protoc -I/usr/local/include -I. \-I$GOPATH/src  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \--swagger_out=logtostderr=true:. \rpc.proto
