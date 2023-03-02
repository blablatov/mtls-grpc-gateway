## Service and Client of gRPC via mTLS. gRPC-gateway

### Building and Running Service

In order to build, Go to ``Go`` module root directory location (mtls-grpc-gateway/gw-mtls-service) and execute the following
 shell command,
```
go build -v 
```  
```
./gw-mtls-service
```

### Building and Running Client   

In order to build, Go to ``Go`` module root directory location (mtls-grpc-gateway/gw-mtls-client) and execute the following
 shell command,
```
go build -v
```  
```
./gw-mtls-client
```

### Testing

* Add a new product to the ProductInfo service.

```
$ curl -X POST http://localhost:8080/v1/product -d '{"name": "Apple", "description": "iphone7", "price": 699}'

"38e13578-d91e-11e9-819f-6c96cfe0687d"
```

* Get the existing product using ProductID

```
$ curl http://localhost:8080/v1/product/38e13578-d91e-11e9-819f-6c96cfe0687d

{"id":"38e13578-d91e-11e9-819f-6c96cfe0687d","name":"Apple","description":"iphone7","price":
```

## Additional Information

### Generate Server and Client side code 
``` 
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. \
product_info.proto
```

### Update after changing the service definition
``` 
go get -u github.com/blablatov/mtls-grpc-gateway/gw-mtls-proto
```

### Generate reverse proxy service code
```
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:. \
product_info.proto
```

### Update after changing the reverse proxy service definition
``` 
go get -u github.com/blablatov/mtls-grpc-gateway/gw-mtls-gate
```

### Generate the swagger file correspond to reverse proxy service
```
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--swagger_out=logtostderr=true:. \
product_info.proto
```

