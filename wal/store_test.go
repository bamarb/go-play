package wal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var data = []byte("Hello World")
var dataLen uint64 = uint64(len(data)) + lenWidth

func TestStoreAppendRead(t *testing.T) {
	f, err := os.CreateTemp("", "store_append_read_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	s, err := newStore(f)
	require.NoError(t, err)
	testAppend(t, s)
	testRead(t, s)
	err = s.Close()
	require.NoError(t, err)
}

func testAppend(t *testing.T, s *store) {
	t.Helper()
	require.NotNil(t, s)
	for i := uint64(1); i < 4; i++ {
		n, pos, err := s.Append(data)
		require.NoError(t, err)
		require.Equal(t, pos+n, dataLen*i)
	}
}

func testRead(t *testing.T, s *store) {
	t.Helper()
	var pos uint64 = 0
	for i := uint64(1); i < 4; i++ {
		b, err := s.Read(pos)
		require.NoError(t, err)
		require.Equal(t, data, b)
		pos += dataLen
	}
}
