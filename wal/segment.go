package wal

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path"
)

type segment struct {
	store      *store
	index      *index
	baseOffset uint64
	nextOffSet uint64
	config     Config
}

func newSegment(dir string, baseOffset uint64, c Config) (*segment, error) {
	storeFile, err := os.OpenFile(
		path.Join(dir, fmt.Sprintf("%d%s", baseOffset, ".store")),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, err
	}
	store, err := newStore(storeFile)
	if err != nil {
		return nil, err
	}

	idxFile, err := os.OpenFile(
		path.Join(dir, fmt.Sprintf("%d%s", baseOffset, ".index")),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, err
	}
	idx, err := newIndex(idxFile, c)
	if err != nil {
		return nil, err
	}
	var noff uint64
	// read the last offset from the index
	if lastOff, _, err := idx.Read(-1); err != nil {
		noff = baseOffset
	} else {
		noff = baseOffset + uint64(lastOff) + 1
	}

	seg := &segment{
		store:      store,
		index:      idx,
		baseOffset: baseOffset,
		nextOffSet: noff,
	}
	return seg, nil
}

// toRelative converts an absolute offset to a relative offset
func (s *segment) toRelative(absOffset uint64) (uint32, error) {
	relOff := absOffset - s.baseOffset
	if relOff < 0 || relOff > math.MaxUint32 {
		return 0, errors.New("invalid offset")
	}
	return uint32(relOff), nil
}

func (s *segment) Append(p []byte) (offset uint64, err error) {
	// write the bytes first to the store
	_, pos, err := s.store.Append(p)
	if err != nil {
		return 0, err
	}
	// write the position to the index file
	// index offsets are relative to the base offset
	relativeOffset, err := s.toRelative(s.nextOffSet)
	if err != nil {
		return 0, err
	}

	err = s.index.Write(relativeOffset, pos)
	if err != nil {
		return 0, err
	}
	cur := s.nextOffSet
	s.nextOffSet++
	return cur, nil
}

func (s *segment) Read(off uint64) ([]byte, error) {
	relOff, err := s.toRelative(off)
	if err != nil {
		return nil, err
	}
	// Lookup the Index for the pos
	_, pos, err := s.index.Read(int64(relOff))
	if err != nil {
		return nil, err
	}
	p, err := s.store.Read(pos)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *segment) IsMaxed() bool {
	return s.store.size >= s.config.Segment.MaxStoreBytes ||
		s.index.size >= s.config.Segment.MaxIndexBytes
}

func (s *segment) Close() error {
	if err := s.index.Close(); err != nil {
		return err
	}
	if err := s.store.Close(); err != nil {
		return err
	}
	return nil
}

func (s *segment) Remove() error {
	if err := s.Close(); err != nil {
		return err
	}
	if err := os.Remove(s.store.Name()); err != nil {
		return err
	}
	if err := os.Remove(s.index.Name()); err != nil {
		return err
	}
	return nil
}

// nearestMultiple returns the nearest lesser multiple of k in j,
// nearestMultiple(9,4) = 8
func nearestMultiple(j, k uint64) uint64 {
	if j > 0 {
		return (j / k) * k
	}
	return ((j - k + 1) / k) * k
}
