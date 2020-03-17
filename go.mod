module github.com/vladfr/arko

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	google.golang.org/grpc v1.28.0
)

replace github.com/vladfr/arko/master/register => ./master/register
