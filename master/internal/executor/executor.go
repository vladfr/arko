package executor

import (
	"context"
	"fmt"
	"os"
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
	Execute(jobName string) (string, error)
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
func (e *JobExecutor) Execute(jobName string) (string, error) {
	slaves := e.db.ActiveSlaves()
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
	defer conn.Close()

	if err != nil {
		return "", fmt.Errorf("Cannot connect to slave %s", slaveAddr)
	}

	// taken from the wonderful grpcurl library
	ctx := context.Background()
	refClient := grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(conn))
	descSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)

	data := ""
	in := strings.NewReader(data)

	rf, formatter, err := grpcurl.RequestParserAndFormatterFor(grpcurl.Format(grpcurl.FormatJSON), descSource, true, false, in)
	h := grpcurl.NewDefaultEventHandler(os.Stdout, descSource, formatter, true)
	err = grpcurl.InvokeRPC(ctx, descSource, conn, jobName, []string{}, h, rf.Next)
	if err != nil {
		return "", fmt.Errorf("Failed to invoke method %s on slave %s", jobName, slaveAddr)
	}
	// end taken from the wonderful grpcurl library
	grpcurl.PrintStatus(os.Stderr, h.Status, formatter)

	return fmt.Sprintf("Job scheduled on slave %s:%d", slave.Config.GetHost(), slave.Config.GetPort()), nil
}
