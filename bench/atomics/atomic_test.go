// go test -cpu="1,2,4,8,16,24" -bench=. -benchtime=100x ./atomics

/*
These set of benchmarks are aimed to test the cost of accessing
shared pointers. There are three tests:

- loopLocalNoPtr: does not increment the shared pointer in the loop.
This is expected to have best performance.

- loopLocal: increments shared pointer in the loop,
but without using atomic package.

- loopAtomic: increments shared pointer in the loop using atomic package.

It is expected that performance of functions accessing shared pointer will
degrade with increased number of goroutines due to CPU cache contention and
invalidation.

These show rapid degradation in performance when shared pointers get modified.
*/
package atomics

import (
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func routineCount() int {
	res := runtime.GOMAXPROCS(-1)
	return res
}

var array = func() []int64 {
	const loopCount = 1000000
	res := make([]int64, loopCount)
	for i := int64(0); i < loopCount; i++ {
		res[i] = i
	}
	return res
}()

// loopLocalNoPtr increments local counter and the total counter without using atomic primitives.
// It increments total counter by taking a local copy first, and adding to it in the end.
func loopLocalNoPtr(array []int64, _ *int64) int64 {
	localCounter := int64(0)
	for _, val := range array {
		localCounter += val
		if localCounter > int64(math.MaxInt64) {
			return localCounter
		}
	}

	return localCounter
}

// loopLocal increments local counter and the total counter without using atomic primitives.
// It increments total counter directly using pointer dereference.
func loopLocal(array []int64, totalCounter *int64) int64 {
	localCounter := int64(0)
	max64 := int64(math.MaxInt64)
	for _, val := range array {
		localCounter += val
		*totalCounter += val
		if localCounter > max64 || *totalCounter > max64 {
			return localCounter
		}
	}

	return localCounter
}

// loopAtomic increments local counter. It uses atomic primitives to increment total counter.
func loopAtomic(array []int64, totalCounter *int64) int64 {
	localCounter := int64(0)
	max64 := int64(math.MaxInt64)
	for _, val := range array {
		localCounter += val
		totalCounterNew := atomic.AddInt64(totalCounter, val)
		if localCounter > max64 || totalCounterNew > max64 {
			return localCounter
		}
	}

	return localCounter
}

func BenchmarkLoopNoPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		routineCount := routineCount()
		for r := 0; r < routineCount; r++ {
			wg.Add(1)
			totalCounter := int64(0)
			go func() {
				defer wg.Done()
				loopLocalNoPtr(array, &totalCounter)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		routineCount := routineCount()
		totalCounter := int64(0)
		for r := 0; r < routineCount; r++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				loopLocal(array, &totalCounter)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkLoopAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		routineCount := routineCount()
		totalCounter := int64(0)
		for r := 0; r < routineCount; r++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				loopAtomic(array, &totalCounter)
			}()
		}
		wg.Wait()
	}
}
