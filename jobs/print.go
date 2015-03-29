package jobs

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type PrintJob struct {
	*baseJob
	message string
}

func NewPrint(message string, wait time.Duration) *PrintJob {
	key := fmt.Sprintf("Print-%s-%d", message, rand.Int31())
	return &PrintJob{
		baseJob: newBaseJob(key, time.Now().Add(wait)),
		message: message,
	}
}

func (job *PrintJob) Run() error {
	log.Printf("%s", job.message)
	return nil
}
