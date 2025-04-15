package wal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTruncate(t *testing.T) {
	f, err := os.OpenFile("/tmp/trunc_test.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	require.NoError(t, err)
	szOpen, err := os.Stat(f.Name())
	require.NoError(t, err)
	require.Equal(t, int64(0), szOpen.Size())
	err = os.Truncate(f.Name(), 1024)
	require.NoError(t, err)
	szTruncate, err := os.Stat(f.Name())
	require.NoError(t, err)
	require.Equal(t, int64(1024), szTruncate.Size())
	os.Remove(f.Name())
}
