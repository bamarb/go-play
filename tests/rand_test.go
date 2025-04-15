package tests

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

const alpanumeric = "abcdefghijklmnopqrstuvwxyz1234567890"

func TestRandRead(t *testing.T) {
	b := make([]byte, 7)
	n, err := rand.Read(b)
	require.NoError(t, err)
	require.Equal(t, 7, n)
	t.Logf("%v", b)
	out := make([]byte, 7)
	for j, rb := range b {
		i := int(rb) % len(alpanumeric)
		out[j] = alpanumeric[i]
	}
	t.Logf("%s", string(out))
}
