### Generate Server and Client side code   
Go to ``Go`` module directory location `stream-mtls-grpc/mtls-proto` and execute the following shell commands: 

```
protoc --go_out=plugins=grpc:. ./product_info.proto
protoc --grpc-gateway_out=logtostderr=true:. ./product_info.proto
```
  
``` 
protoc product_info.proto --go_out=./ --go-grpc_out=./
protoc product_info.proto --go_out=./ --grpc-gateway_out=./
protoc product_info.proto --grpc-gateway_out

protoc --grpc-gateway_out=. .product_info.proto
``` 

```
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:. \
product_info.proto
```

```
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:. \
product_info.proto
```