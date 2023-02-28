module github.com/blablatov/mtls-grpc-gateway

go 1.18

replace github.com/blablatov/mtls-grpc-gateway/gw-mtls-proto => ./gw-mtls-proto

replace github.com/blablatov/mtls-grpc-gateway/gw-mockups => ./gw-mockups

require (
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	golang.org/x/oauth2 v0.4.0
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	cloud.google.com/go/compute v1.15.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
