module github.com/vladfr/arko

go 1.13

require (
	github.com/asdine/storm v2.1.2+incompatible
	github.com/asdine/storm/v3 v3.1.1
	github.com/favadi/protoc-go-inject-tag v1.0.0 // indirect
	github.com/fullstorydev/grpcurl v1.5.0
	github.com/golang/protobuf v1.3.5
	github.com/jhump/protoreflect v1.5.0
	google.golang.org/grpc v1.28.0
	istio.io/istio v0.0.0-20200330071856-c62e39664812
)

replace github.com/vladfr/arko/master/register => ./master/register
