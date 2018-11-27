ifndef GOOS
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	GOOS := darwin
else ifeq ($(UNAME_S),Linux)
	GOOS := linux
else
$(error "$$GOOS is not defined. If you are using Windows, try to re-make using 'GOOS=windows make ...' ")
endif
endif

# dev test online
PROJECTENV=dev

#BUILD_FLAGS := -ldflags "-X v1.0"

env:
	@echo "Building project env"
	export GOENV=$(PROJECTENV)

rpc:
	@echo "Building golang grpc server"
	@cd rpc

	@echo "Building google.api"
	protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

	@echo "Building rpc.pb.proto"
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc,Mgoogle/api/annotations.proto=wps_store/rpc/google/api:. \
	rpc.proto

	@echo "Building rpc.pb.gw.proto"
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:. \
	rpc.proto

	@echo "Building swagger.json"
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:. \
	rpc.proto

build : server http-server

server:
	@echo "Building micro store to cmd/server"

http-server:
	@echo "Building micro to cmd/server_http"

clean:
	@echo "Cleaning binaries built"
	@rm -f cmd/server
	@rm -f cmd/server_http
