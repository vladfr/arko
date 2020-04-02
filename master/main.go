package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"

	execpb "github.com/vladfr/arko/master/execution"
	"github.com/vladfr/arko/master/internal/executor"
	"github.com/vladfr/arko/master/models"
	pb "github.com/vladfr/arko/master/register"
)

// SlavePingSeconds sets the interval to ping slaves
const SlavePingSeconds = 30

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	dbFile   = flag.String("db", "arko.db", "The Storm/bbolt database file")
	port     = flag.Int("port", 10001, "The server port")
)

func (s *registerServer) pingSlaves(ticker *time.Ticker, done chan bool) {
	for {
		select {
		case <-done:
			// stop channel, we should stop
			return
		case <-ticker.C:
			// ticked, we should run
			fmt.Println("==== Pinging slaves...")
			for _, slave := range s.db.ActiveSlaves() {
				slaveAddr := fmt.Sprintf("%s:%d", slave.Config.GetHost(), slave.Config.GetPort())
				conn, err := grpc.Dial(slaveAddr,
					grpc.WithInsecure(),
					grpc.WithBlock(),
					grpc.WithTimeout(5*time.Second),
				)

				if err != nil {
					fmt.Println("Cannot connect to slave", slaveAddr)
				} else {
					fmt.Println("Opened connection to", slaveAddr)
					fmt.Println("Slave ", slaveAddr, "has status ", slave.Status)
					fmt.Println("Updating methods of ", slaveAddr)
					methods, _ := s.reflectOnSlave(conn)
					if len(methods) > 0 {
						slave.Methods = methods
						fmt.Errorf("\tCould not find any methods on slave")
					}
					slave.Status = 1
					s.db.SaveSlave(&slave)
					fmt.Println("\tMethods on ", slaveAddr, "are :", slave.Methods)
					conn.Close()
				}
			}
			fmt.Println("==== Pinging slaves done")
		}
	}
}

func (s *registerServer) reflectOnSlave(conn *grpc.ClientConn) (methods []string, err error) {
	ctx := context.Background()
	refClient := grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(conn))
	descSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)
	svcs, err := grpcurl.ListServices(descSource)
	if err != nil {
		fmt.Errorf("\tFailed to list services: %v", err)
	}
	if len(svcs) == 0 {
		fmt.Println("\t(No services)")
	} else {
		// Code taken from the wonderful grpcurl library
		for _, svc := range svcs {
			sname := fmt.Sprintf("%s\n", svc)
			// dsc, err = descSource.FindSymbol(svc)
			// if err != nil {
			// 	fmt.Println(err, "Failed to resolve symbol", sname)
			// }
			// fmt.Println(dsc.GetFullyQualifiedName())
			svcMethods, err := grpcurl.ListMethods(descSource, svc)
			methods = append(methods, svcMethods...)
			if err != nil {
				fmt.Println("\tFailed to list methods for service", sname, err)
			} else if len(methods) == 0 {
				fmt.Println("\t(No methods found)")
			}
		}
	}
	return
}

type registerServer struct {
	pb.UnimplementedRegisterServer
	db models.Datastore
}

type executionServer struct {
	execpb.UnimplementedExecutionServer
	db   models.Datastore
	exec executor.Executor
}

func (s *registerServer) RegisterNewSlave(ctx context.Context, config *pb.SlaveConfig) (*pb.SlaveRegisterStatus, error) {
	s.db.AddSlave(config)
	fmt.Printf("Slave registered at %v:%d", config.GetHost(), config.GetPort())
	fmt.Println()
	return &pb.SlaveRegisterStatus{Message: "done"}, nil
}

func (s *executionServer) ExecuteJob(ctx context.Context, params *execpb.JobParams) (*execpb.JobStatus, error) {
	paramsJSON, _ := json.Marshal(params.GetParams())
	result, err := s.exec.Execute(params.Method, string(paramsJSON))
	if err != nil {
		result = err.Error()
	}
	msg := fmt.Sprintf("%s done: %s", params.Method, result)
	return &execpb.JobStatus{Message: msg}, nil
}

func main() {
	fmt.Println("Master starting...")
	flag.Parse()

	db, err := models.NewDB(*dbFile)
	if err != nil {
		log.Panicf("Cannot load/create database file %s", dbFile)
	}

	exec := executor.NewJobExecutor(db)

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

	registerServer := &registerServer{db: db}
	pb.RegisterRegisterServer(grpcServer, registerServer)
	execpb.RegisterExecutionServer(grpcServer, &executionServer{db: db, exec: exec})
	fmt.Println("Server listening for slaves on port", *port)

	ticker := time.NewTicker(time.Duration(SlavePingSeconds) * time.Second)
	done := make(chan bool)
	go registerServer.pingSlaves(ticker, done)

	grpcServer.Serve(lis)
}
