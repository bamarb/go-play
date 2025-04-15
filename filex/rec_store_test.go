package filex

import (
	"encoding/binary"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	_ io.Reader = (*rec)(nil)
	_ io.Writer = (*rec)(nil)
)

type rec struct {
	kind uint32
	data string
}

func (r rec) Read(p []byte) (n int, err error) {
	binary.BigEndian.PutUint32(p, r.kind)
	n += 4
	binary.BigEndian.PutUint32(p[n:], uint32(len(r.data)))
	n += 4
	for i := 0; i < len(r.data); i++ {
		p[i+n] = r.data[i]
	}
	n += len(r.data)
	return n, nil
}

func (r *rec) Write(p []byte) (n int, err error) {
	kind := binary.BigEndian.Uint32(p[0:4])
	dlen := binary.BigEndian.Uint32(p[4:8])
	fmt.Printf("kind: %d , dlen: %d \n", kind, dlen)
	dbytes := make([]byte, dlen)
	copy(dbytes, p[8:])
	r.kind = kind
	r.data = string(dbytes)
	return
}

func TestReadWrite(t *testing.T) {
	r := &rec{kind: 10, data: "Hallelujah"}
	p := make([]byte, 64)
	n, err := r.Read(p)
	require.NoError(t, err)
	require.Equal(t, 4+4+len(r.data), n)
	trec := &rec{}
	trec.Write(p)
	assert.Equal(t, r, trec)
}
