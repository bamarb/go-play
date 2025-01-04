package bytex_test

import (
	"bytes"
	"encoding/binary"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestByteBufferBasic(t *testing.T) {
	var bbuf bytes.Buffer
	t.Logf("bbuf len: %d cap:%d avail:%d", bbuf.Len(), bbuf.Cap(), bbuf.Available())
	ip := [4]byte{15, 16, 17, 18}
	for i := 0; i < len(ip); i++ {
		t.Logf("%d", ip[i])
	}
	bbuf.Write(ip[:])
	if bbuf.Len() != 4 {
		t.Errorf("error got %d want %d", bbuf.Len(), 4)
	}
	t.Logf("bbuf len: %d cap:%d avail:%d", bbuf.Len(), bbuf.Cap(), bbuf.Available())

	sip := make([]byte, 4)
	bbuf.Read(sip)
	t.Logf("%v", sip)
	t.Logf("bbuf len: %d cap:%d avail:%d", bbuf.Len(), bbuf.Cap(), bbuf.Available())
	//	sip = slices.Replace(sip, 0, len(sip), 0, 0, 0, 0)
	clear(sip)
	t.Logf("len: %d , cap: %d", len(sip), cap(sip))
	t.Logf("%v", sip)
	_, err := bbuf.Read(sip)
	t.Logf("%v", sip)
	t.Logf("bbuf len: %d cap:%d avail:%d", bbuf.Len(), bbuf.Cap(), bbuf.Available())
	if err != nil {
		t.Logf("%s", err)
	}
}

func TestByteBufferWithFile(t *testing.T) {
	file := "/tmp/gobbuf.tmp"
	fh, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0666))
	require.NoError(t, err)
	x := 0xdeadbeef
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, x)
}
