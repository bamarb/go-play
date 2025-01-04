package tests

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestTimeTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(2 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
	LOOP:
		for {
			select {
			case <-ticker.C:
				t.Logf("%s :-> tick", time.Now().String())
			case <-ctx.Done():
				break LOOP
			}
		}
		t.Logf("Ticker done...")
		wg.Done()
	}()
	time.Sleep(12 * time.Second)
	ticker.Stop()
	cancel()
	wg.Wait()
}

func TestCancelMult(t *testing.T) {
	rootCtx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(ctx context.Context, num int) {
			timerCh := time.NewTicker(1 * time.Second)
		LOOP:
			for {
				select {
				case <-ctx.Done():
					break LOOP
				case <-timerCh.C:
					t.Logf("timer fired: GR : %d\n", num)
				}
			}
			t.Logf("Go routine %d done\n", num)
			wg.Done()
		}(rootCtx, i)
	}
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}
