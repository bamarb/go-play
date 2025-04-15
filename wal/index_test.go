package wal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "index_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	c := Config{}
	c.Segment.MaxIndexBytes = 1 << 10
	idx, err := newIndex(f, c)
	require.NoError(t, err)
	_, _, err = idx.Read(-1)
	require.Error(t, err)
	entries := []struct {
		off uint32
		pos uint64
	}{
		{0, 0},
		{1, 10},
		{2, 20},
	}
	for _, want := range entries {
		err := idx.Write(want.off, want.pos)
		require.NoError(t, err)
		o, p, err := idx.Read(int64(want.off))
		require.NoError(t, err)
		require.Equal(t, want.pos, p)
		require.Equal(t, want.off, o)
	}
	// index and scanner should error when reading past entries
	_, _, err = idx.Read(int64(len(entries)))
	require.Error(t, err)
	_ = idx.Close()
	// Index Should rebuild it's entries from existing file
	idxf, err := os.OpenFile(f.Name(), os.O_RDWR, 0600)
	require.NoError(t, err)
	ridx, err := newIndex(idxf, c)
	require.NoError(t, err)
	off, pos, err := ridx.Read(-1)
	require.NoError(t, err)
	require.Equal(t, entries[len(entries)-1].off, off)
	require.Equal(t, entries[len(entries)-1].pos, pos)
	_ = ridx.Close()
}
