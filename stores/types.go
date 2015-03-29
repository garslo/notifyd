package stores

import (
	"time"

	"github.com/garslo/notifyd/jobs"
)

type Store interface {
	Add(jobs.Job) error
	Get(name string) (jobs.Job, error)
	Cancel(name string) error
	GetUpcoming(time.Duration) (jobs.Jobs, error)
	List() (jobs.Jobs, error)
}
