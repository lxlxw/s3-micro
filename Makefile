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

PROJECTNAME=s3-micro

all: build
	
grpc:
	@echo "Building golang grpc server"

	@echo "Building rpc.pb.proto"
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc,Mgoogle/api/annotations.proto=$(PROJECTNAME)/proto/google/api:. \
	./proto/rpc.proto

	@echo "Building rpc.pb.gw.proto"
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:. \
	./proto/rpc.proto

	@echo "Building swagger.json"
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:. \
	./proto/rpc.proto

build : 
	@echo "Building s3-micro"
	@go build

server:
	@echo "Running s3-micro to cmd/server"
	@./$(PROJECTNAME) server

http:
	@echo "Running s3-micro to cmd/server_http"
	@./$(PROJECTNAME) http

clean:
	@echo "Cleaning binaries built"
	@rm -f $(PROJECTNAME)
	@rm -f proto/google/api/annotations.pb.go
	@rm -f proto/google/api/http.pb.go
	@rm -f proto/rpc.pb.go
	@rm -f proto/rpc.pb.gw.go
	@rm -f proto/rpc.swagger.json

swagger:
	@echo "Building pkg/swagger/swagger.go"
	@go-bindata --nocompress -pkg swagger -o pkg/swagger/swagger.go swagger-ui/

