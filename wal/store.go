package wal

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

// Record - the data that we store in the log
// Store - the file we store records in
// Index - the file we store index entries in
// Segment - An abstraction that ties a store and index together
// Log - the abstraction that ties all the segments together

var enc = binary.BigEndian

const lenWidth = 8
const storeSuffix = ".store"

type store struct {
	*os.File
	mu   sync.Mutex
	buf  *bufio.Writer
	size uint64
}

func newStore(f *os.File) (*store, error) {
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	// get the current size of the file in case we are creating
	// a store from an already existing file
	sz := uint64(fi.Size())
	return &store{File: f, size: sz, buf: bufio.NewWriter(f)}, nil
}

func (s *store) Append(data []byte) (n uint64, pos uint64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	pos = s.size
	if err = binary.Write(s.buf, enc, uint64(len(data))); err != nil {
		return 0, 0, err
	}
	w, err := s.buf.Write(data)
	if err != nil {
		return 0, 0, err
	}
	w += lenWidth // Add 8 Bytes (uint64) for recording the len
	s.size += uint64(w)
	return uint64(w), pos, nil
}

// Read returns the data at the provided position excluding the length encoding
func (s *store) Read(pos uint64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.buf.Flush(); err != nil {
		return nil, err
	}
	size := make([]byte, lenWidth)
	if _, err := s.File.ReadAt(size, int64(pos)); err != nil {
		return nil, err
	}
	data := make([]byte, enc.Uint64(size))
	if _, err := s.File.ReadAt(data, int64(pos+lenWidth)); err != nil {
		return nil, err
	}
	return data, nil
}

// ReadAt reads len(p) bytes into p starting at the offset (off)
func (s *store) ReadAt(p []byte, off int64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.buf.Flush(); err != nil {
		return 0, err
	}
	return s.File.ReadAt(p, off)
}

func (s *store) Close() error {
	s.mu.Lock()
	err := s.buf.Flush()
	if err != nil {
		return err
	}
	return s.File.Close()
}
