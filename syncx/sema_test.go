package syncx

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"
)

func TestSemaphoreBasic(t *testing.T) {
	semx := semaphore.NewWeighted(10)
	dones := []chan struct{}{}
	for i := range 10 {
		d := make(chan struct{})
		dones = append(dones, d)
		go func(sem *semaphore.Weighted, id int, done chan struct{}) {
			if semx.TryAcquire(1) {
				fmt.Printf("go routine id: %d accquire \n", i)
				time.Sleep(1 * time.Second)
			}
			semx.Release(1)
			fmt.Printf("go routine id: %d release \n", i)
			close(done)
		}(semx, i, d)
	}

	for _, c := range dones {
		<-c
	}
}
