module github.com/blablatov/mtls-grpc-gateway

go 1.18

replace github.com/blablatov/mtls-grpc-gateway/gw-mtls-proto => ./gw-mtls-proto

replace github.com/blablatov/mtls-grpc-gateway/gw-mockups => ./gw-mockups

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	google.golang.org/genproto v0.0.0-20230223222841-637eb2293923 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
