# wps store micro service


# CA TLS
cd conf/certs
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

Country Name (2 letter code) [XX]:
State or Province Name (full name) []:
Locality Name (eg, city) [Default City]:
Organization Name (eg, company) [Default Company Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (eg, your name or your server's hostname) []:wps grpc micro service
Email Address []:

# protoc-gen
mv protoc-gen-swagger /usr/local/go/bin
mv protoc-gen-grpc-gateway /usr/local/go/bin


# 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

#编译rpc.proto为rpc.pb.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=wps_store/proto/google/api:. ./rpc.proto

#编译rpc.proto为rpc.pb.gw.proto
protoc --grpc-gateway_out=logtostderr=true:. ./rpc.proto

# TODO：
# 1.TLS认证
# 2.JWT TOKEN