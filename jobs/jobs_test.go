package jobs

import (
	"testing"
	"time"
)

func TestThatJobsAreRemoved(t *testing.T) {
	js := Jobs{}
	p1 := NewPrint("foo", 1*time.Second)
	js.Add(p1)
	js.Remove(p1.Name())
	if len(js) != 0 {
		t.Errorf("len(js) != 0: js = %v", js)
	}
}
