package jobs

import "time"

type Job interface {
	Name() string
	ScheduledFor() time.Time
	Run() error
}

type Jobs []Job

func (jobs Jobs) Get(name string) (Job, error) {
	for _, job := range jobs {
		if job.Name() == name {
			return job, nil
		}
	}
	return nil, ErrNoSuchJob
}

func (jobs Jobs) Add(job Job) {
	jobs = append(jobs, job)
}

func (jobs *Jobs) Remove(name string) {
	for i, job := range *jobs {
		if job.Name() == name {
			*jobs = append((*jobs)[:i], (*jobs)[i+1:]...)
			return
		}
	}
}
