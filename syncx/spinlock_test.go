package syncx

import (
	"sync"
	"testing"
)

func TestSimpleCounterLock(t *testing.T) {
	lock := &SpinLock{}
	counter := 0
	iter := 100000
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		lock.Lock()
		for i := 0; i < iter; i++ {
			counter += 1
		}
		lock.Unlock()
		wg.Done()
	}()
	go func() {
		lock.Lock()
		for i := 0; i < iter; i++ {
			counter += 1
		}
		lock.Unlock()
		wg.Done()
	}()
	wg.Wait()
	t.Logf("Counter Value: %d\n", counter)
}
