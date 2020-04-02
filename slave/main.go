package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pbm "github.com/vladfr/arko/master/register"
	pb "github.com/vladfr/arko/slave/pipeline"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	serverAddr = flag.String("master", "127.0.0.1:10001", "Master to connect to")
	port       = flag.Int("port", 10002, "The server port")
)

// GetOutboundIP Gets preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

type pipeline struct {
	pb.UnimplementedMyPipelineServer
}

func (s *pipeline) Run(ctx context.Context, config *pb.MyPipelineConfig) (*pb.PipelineStatus, error) {
	time.Sleep(5 * time.Second)
	fmt.Println("Executed pipeline MyPipeline.Run with ", config.GetParam(), config.GetPassword())
	return &pb.PipelineStatus{Message: fmt.Sprintf("job execution done with: %s %s", config.GetParam(), config.GetPassword())}, nil
}

func (s *pipeline) DryRun(ctx context.Context, config *pb.MyPipelineConfig) (*pb.PipelineStatus, error) {
	fmt.Println("Executed pipeline with DryRun")
	return &pb.PipelineStatus{Message: "job execution done"}, nil
}

func (s *pipeline) Rollback(ctx context.Context, config *pb.MyPipelineConfig) (*pb.PipelineStatus, error) {
	fmt.Println("Executed rollback on pipeline")
	return &pb.PipelineStatus{Message: "job execution done"}, nil
}

func startSlave() {
	fmt.Println("Starting slave...")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	// if *tls {
	// 	if *certFile == "" {
	// 		*certFile = testdata.Path("server1.pem")
	// 	}
	// 	if *keyFile == "" {
	// 		*keyFile = testdata.Path("server1.key")
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate credentials %v", err)
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMyPipelineServer(grpcServer, &pipeline{})
	reflection.Register(grpcServer)
	fmt.Println("Slave listening for connections on port", *port)

	grpcServer.Serve(lis)
}

func main() {

	go startSlave()

	fmt.Println("Registering on master", *serverAddr)
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := pbm.NewRegisterClient(conn)
	res, err := client.RegisterNewSlave(context.Background(), &pbm.SlaveConfig{
		Host:  fmt.Sprintf("%s", GetOutboundIP()),
		Port:  int32(*port),
		Token: "asdasd",
	})
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Registered on master", *serverAddr)
	}

	select {}
}
