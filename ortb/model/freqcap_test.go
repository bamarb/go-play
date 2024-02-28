package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFreqCap_zero(t *testing.T) {
	var fcaput Fcap

	res := fcaput.Mark(time.Now())
	require.True(t, res)

	for i := 1; i < 10; i++ {
		tis := time.Now().Add(time.Duration(i) * time.Hour)
		res := fcaput.Mark(tis)
		require.False(t, res)
	}
}

func TestFreqCap_minutes(t *testing.T) {
	// 1 Impression Every 1 Minutes
	spec := FreqencyCap{MaxImpressions: 1, NumTimeUnits: 1, TimeUnit: MINUTE}
	fcap := NewFcapOf(spec)
	reft := time.Now()
	require.True(t, fcap.Mark(reft))
	reft = reft.Add(10 * time.Second)
	require.False(t, fcap.Mark(reft))
	reft = reft.Add(10 * time.Second)
	require.False(t, fcap.Mark(reft))
	reft = reft.Add(40 * time.Second)
	ret := fcap.Mark(reft)
	require.True(t, ret)
	require.False(t, fcap.Mark(reft.Add(1*time.Second)))
}

func TestFreqCap_lifetime(t *testing.T) {
	// 1 Impression forever which is same as the zero value
	t.Run("1 Forever", func(t *testing.T) {
		spec := FreqencyCap{MaxImpressions: 1, NumTimeUnits: 1, TimeUnit: LIFETIME}
		fcap := NewFcapOf(spec)
		reft := time.Now()
		require.True(t, fcap.Mark(reft))
		for i := 1; i < 366; i++ {
			tt := reft.Add(time.Duration(i) * 24 * time.Hour)
			require.False(t, fcap.Mark(tt))

		}
	})
	t.Run("2 Forever", func(t *testing.T) {
		spec := FreqencyCap{MaxImpressions: 2, NumTimeUnits: 1, TimeUnit: LIFETIME}
		fcap := NewFcapOf(spec)
		reft := time.Now()
		require.True(t, fcap.Mark(reft))
		reft = reft.Add(time.Duration(30) * 24 * time.Hour)
		require.True(t, fcap.Mark(reft))
		for i := 1; i < 366; i++ {
			tt := reft.Add(time.Duration(i) * 24 * time.Hour)
			require.False(t, fcap.Mark(tt))

		}
	})
}
