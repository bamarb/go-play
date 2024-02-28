package randx

import (
	"testing"
)

func TestIntBetween(t *testing.T) {
	for i := 0; i < 2000; i++ {
		val := IntBetween(50, 100)
		// t.Logf("rand num %d\n", val)
		if val < 50 || val > 99 {
			t.FailNow()
		}
	}
}

func TestCoinToss(t *testing.T) {
	const Heads = true
	const Tails = false

	for i := 0; i < 2000; i++ {
		res := CoinToss()
		if res != Heads {
			if res != Tails {
				t.FailNow()
			}
		}
	}
}
