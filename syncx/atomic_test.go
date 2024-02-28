package syncx

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Data struct {
	id      string
	counter atomic.Uint64
}

func (d *Data) Increment() {
	d.counter.Add(1)
}

func DataMaker(idpfx string) []*Data {
	ret := []*Data{}
	for i := 0; i < 1024; i++ {
		id := fmt.Sprintf("%s-%d", idpfx, i)
		d := &Data{id: id}
		ret = append(ret, d)
	}
	return ret
}

type DataSet struct {
	Id      string
	dataPtr atomic.Pointer[[]*Data]
}

func New(pfx string) *DataSet {
	d := DataMaker(pfx)
	ret := &DataSet{Id: pfx}
	ret.dataPtr.Store(&d)
	return ret
}

func TestAtomicPointer(t *testing.T) {
	n := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(n)
	// start go routines
	dset := New("data1")
	go func() {
		time.Sleep(1 * time.Second)
		dset2 := DataMaker("data2")
		dset.dataPtr.Store(&dset2)
		t.Logf("Changed the ptr %p\n", &dset2)
	}()

	for i := 0; i < n; i++ {
		go func(ds *DataSet, i int) {
			dslicePtr := ds.dataPtr.Load()
			dslice := *dslicePtr
			for j := 0; j < 3; j++ {
				time.Sleep(1 * time.Second)
				t.Logf("routine %d operating on ptr %p\n", i, dslicePtr)
				mydata := dslice[i]
				mydata.Increment()
			}
			wg.Done()
		}(dset, i)
	}
	wg.Wait()
}
