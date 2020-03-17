package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/vladfr/arko/master/register"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10001, "The server port")
)

func pingSlaves(ticker *time.Ticker, done chan bool) {
	for {
		select {
		case <-done:
			// stop channel, we should stop
			return
		case <-ticker.C:
			// ticked, we should run
			fmt.Println("Pinging slaves...")
			for _, s := range slaveList {
				slaveAddr := fmt.Sprintf("%s:%d", s.config.GetHost(), s.config.GetPort())
				conn, err := grpc.Dial(slaveAddr,
					grpc.WithInsecure(),
					grpc.WithBlock(),
					grpc.WithTimeout(5*time.Second),
				)

				if err != nil {
					fmt.Println("Cannot connect to slave", slaveAddr)
				} else {
					fmt.Println("Opened connection to", slaveAddr)
				}
				conn.Close()
			}
			fmt.Println("Pinging slaves done")
		}
	}
}

type Slave struct {
	config *pb.SlaveConfig
	conn   *grpc.ClientConn
}

var slaveList []*Slave

func NewSlave(config *pb.SlaveConfig) *Slave {
	return &Slave{
		config: config,
		conn:   nil,
	}
}

type registerServer struct {
}

func (s *registerServer) RegisterNewSlave(ctx context.Context, config *pb.SlaveConfig) (*pb.SlaveRegisterStatus, error) {
	slaveList = append(slaveList, NewSlave(config))
	fmt.Printf("Slave registered at %v:%d", config.GetHost(), config.GetPort())
	fmt.Println()
	return &pb.SlaveRegisterStatus{Message: "done"}, nil
}

func main() {
	fmt.Println("Master starting...")
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
	pb.RegisterRegisterServer(grpcServer, &registerServer{})
	fmt.Println("Server listening for slaves on port", *port)

	ticker := time.NewTicker(time.Duration(5) * time.Second)
	done := make(chan bool)
	go pingSlaves(ticker, done)

	grpcServer.Serve(lis)
}
