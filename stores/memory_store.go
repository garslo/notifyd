package stores

import (
	"time"

	"github.com/garslo/notifyd/jobs"
)

type MemoryStore struct {
	jobList jobs.Jobs
}

func NewMemory() *MemoryStore {
	return &MemoryStore{
		jobList: jobs.Jobs{},
	}
}

func (store *MemoryStore) Add(job jobs.Job) error {
	store.jobList = append(store.jobList, job)
	return nil
}

func (store *MemoryStore) Get(name string) (jobs.Job, error) {
	return store.jobList.Get(name)
}

func (store *MemoryStore) Cancel(name string) error {
	store.jobList.Remove(name)
	return nil
}

func (store *MemoryStore) GetUpcoming(dur time.Duration) (jobs.Jobs, error) {
	upcoming := jobs.Jobs{}
	now := time.Now()
	for _, job := range store.jobList {
		if job.ScheduledFor().Sub(now) <= dur {
			upcoming = append(upcoming, job)
			store.jobList.Remove(job.Name())
		}
	}
	return upcoming, nil
}

func (store *MemoryStore) List() (jobs.Jobs, error) {
	return store.jobList, nil
}
