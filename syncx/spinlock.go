package syncx

import (
	"runtime"
	"sync/atomic"
)

// SpinLock implements Locker default value false (unlocked)
type SpinLock struct {
	lock atomic.Bool
}

func (sl *SpinLock) Lock() {
	for !sl.TryLock() {
		runtime.Gosched()
	}
}

func (sl *SpinLock) Unlock() {
	sl.lock.Store(false)
}

func (sl *SpinLock) TryLock() bool {
	return sl.lock.CompareAndSwap(false, true)
}
