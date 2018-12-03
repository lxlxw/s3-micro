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

PROJECTNAME=micro-s3-grpc

all: build
	
grpc:
	@echo "Building golang grpc server"
	@cd rpc

	@echo "Building rpc.pb.proto"
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc,Mgoogle/api/annotations.proto=$(PROJECTNAME)/rpc/google/api:. \
	./rpc/rpc.proto

	@echo "Building rpc.pb.gw.proto"
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:. \
	./rpc/rpc.proto

	@echo "Building swagger.json"
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:. \
	./rpc/rpc.proto

	@echo "Success!"

build : 
	@echo "Building micro store"
	@go build

server:
	@echo "Running micro store to cmd/server"
	@./$(PROJECTNAME) server

http:
	@echo "Running micro store to cmd/server_http"
	@./$(PROJECTNAME) http

clean:
	@echo "Cleaning binaries built"
	@rm -f $(PROJECTNAME)
	@rm -f rpc/google/api/annotations.pb.go
	@rm -f rpc/google/api/http.pb.go
	@rm -f rpc/rpc.pb.go
	@rm -f rpc/rpc.pb.gw.go
	@rm -f rpc/rpc.swagger.json

swagger:
	@echo "Building pkg/swagger/swagger.go"
	@go-bindata --nocompress -pkg swagger -o pkg/swagger/swagger.go swagger-ui/

