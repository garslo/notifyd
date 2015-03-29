package runners

import (
	"log"
	"sync"
	"time"

	"github.com/garslo/notifyd/jobs"
	"github.com/garslo/notifyd/stores"
)

type runner struct {
	store       stores.Store
	granularity time.Duration
}

func New(store stores.Store, granularity time.Duration) Runner {
	return &runner{
		store:       store,
		granularity: granularity,
	}
}

func (run *runner) RunNext() []error {
	js, err := run.store.GetUpcoming(run.granularity)
	if err != nil {
		return []error{err}
	}
	wg := &sync.WaitGroup{}
	for _, j := range js {
		wg.Add(1)
		go func(job jobs.Job) {
			<-time.After(job.ScheduledFor().Sub(time.Now()))
			err := job.Run()
			log.Printf("Ran job name='%s' scheduled_for='%s' err='%v'", job.Name(), job.ScheduledFor(), err)
			wg.Done()
		}(j)
	}
	wg.Wait()
	return nil
}

func (run *runner) RunForever() {
	for {
		run.RunNext()
		<-time.Tick(run.granularity)
	}
}
