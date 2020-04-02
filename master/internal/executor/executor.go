package executor

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/vladfr/arko/master/internal/scheduler"
	"github.com/vladfr/arko/master/models"
	"google.golang.org/grpc"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type Executor interface {
	Execute(jobName string, dataJSON string) (string, error)
}

type JobExecutor struct {
	db        models.Datastore
	scheduler scheduler.Scheduler
}

func NewJobExecutor(db models.Datastore) *JobExecutor {
	sc := scheduler.NewDefaultScheduler()
	return &JobExecutor{
		db:        db,
		scheduler: sc,
	}

}

// Execute runs the job and returns the results
func (e *JobExecutor) Execute(jobName string, dataJSON string) (string, error) {
	slaves := e.db.GetActiveSlaves()
	var jobList []string

	slave, err := e.scheduler.FindSlaveForJob(jobName, jobList, slaves)
	fmt.Println(slave, err)
	if err != nil {
		return "", err
	}

	slaveAddr := fmt.Sprintf("%s:%d", slave.Config.GetHost(), slave.Config.GetPort())
	conn, err := grpc.Dial(slaveAddr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)

	if err != nil {
		fmt.Printf("Cannot connect to slave %s", slaveAddr)
		return "", fmt.Errorf("Cannot connect to slave %s", slaveAddr)
	}

	defer conn.Close()

	// taken from the wonderful grpcurl library
	ctx := context.Background()
	refClient := grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(conn))
	descSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)
	// dsc, _ := descSource.FindSymbol(jobName)
	// msg := dsc.GetOptions()
	// msg.ProtoMessage() // get the protoreflect message

	in := strings.NewReader(dataJSON)

	rf, formatter, err := grpcurl.RequestParserAndFormatterFor(grpcurl.Format(grpcurl.FormatJSON), descSource, true, false, in)
	var buf bytes.Buffer

	h := NewBufferedEventHandler(&buf, descSource, formatter)
	err = grpcurl.InvokeRPC(ctx, descSource, conn, jobName, []string{}, h, rf.Next)
	if err != nil {
		return "", fmt.Errorf("Failed to invoke method %s on slave %s: %s", jobName, slaveAddr, err.Error())
	}
	//grpcurl.PrintStatus(os.Stderr, h.Status, formatter)
	// end taken from the wonderful grpcurl library

	fmt.Printf("Job scheduled on slave %s:%d", slave.Config.GetHost(), slave.Config.GetPort())
	return fmt.Sprintf("Job finished with status %s: %s", h.Status.Code(), string(buf.Bytes())), nil
}
