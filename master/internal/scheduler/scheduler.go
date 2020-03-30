package scheduler

import (
	"fmt"

	"github.com/vladfr/arko/master/models"
)

type Scheduler interface {
	FindSlaveForJob(jobToRun string, jobList []string, slaves []models.Slave) (*models.Slave, error)
}

type DefaultScheduler struct {
	jobList []string
	slaves  []*models.Slave
}

// NewDefaultScheduler creates and returns an instance or our default scheduler
func NewDefaultScheduler() *DefaultScheduler {
	return &DefaultScheduler{}
}

// FindSlaveForJob interates through all the slaves and returns
// the first available slave which implements our job
func (sc *DefaultScheduler) FindSlaveForJob(jobToRun string, jobList []string, slaves []models.Slave) (*models.Slave, error) {
	for _, slave := range slaves {
		if slave.Status == 1 {
			for _, method := range slave.Methods {
				if method == jobToRun {
					return &slave, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("Cannot find a suitable slave to schedule %s", jobToRun)
}
