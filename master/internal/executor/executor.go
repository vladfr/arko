package executor

import (
	"fmt"

	"github.com/vladfr/arko/master/internal/scheduler"
	"github.com/vladfr/arko/master/models"
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
	return fmt.Sprintf("Job scheduled on slave %s:%d", slave.Config.GetHost(), slave.Config.GetPort()), nil
}
