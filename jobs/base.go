package jobs

import "time"

type baseJob struct {
	name         string
	scheduledFor time.Time
}

func newBaseJob(name string, scheduledFor time.Time) *baseJob {
	return &baseJob{
		name:         name,
		scheduledFor: scheduledFor,
	}
}

func (base *baseJob) Name() string {
	return base.name
}

func (base *baseJob) ScheduledFor() time.Time {
	return base.scheduledFor
}
