package wal

type Record struct {
	Offset uint64
	Value  []byte
}
