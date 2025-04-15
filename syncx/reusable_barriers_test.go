package syncx

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	loopGate sync.WaitGroup
	bootGate sync.WaitGroup
)

type counter struct {
	sync.Mutex
	c int
}

func (c *counter) Inc() {
	c.Lock()
	c.c += 1
	c.Unlock()
}

func (c *counter) Get() int {
	var res int
	c.Lock()
	res = c.c
	c.Unlock()
	return res
}

func TestCounter(t *testing.T) {
	ctr := &counter{}
	wgStart := &sync.WaitGroup{}
	wgStart.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			for range 10 {
				ctr.Inc()
			}
			wgStart.Done()
		}()
	}
	wgStart.Wait()
	require.Equal(t, 30, ctr.Get())
}

func job(c *counter) {
	time.Sleep(time.Duration(rand.IntN(100)) * time.Millisecond)
	fmt.Printf("counter: %d\n", c.Get())
}

// worker does a series of jobs
func worker(c *counter, id int) {
	fmt.Printf("launched worker: %d\n", id)
	for i := 0; i < 3; i++ {
		c.Inc()
	}
	loopGate.Done()
}

func TestBarrierRaw(t *testing.T) {
	nWorkers := 3
	loopGate.Add(nWorkers)
	ctr := &counter{}
	for i := 0; i < nWorkers; i++ {
		go worker(ctr, i)
	}
	loopGate.Wait()
}
