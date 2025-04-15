package wal

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Log struct {
	mu        sync.Mutex
	Dir       string
	Cfg       Config
	activeSeg *segment
	segments  []*segment
}

func NewLog(dir string, cfg Config) (*Log, error) {
	l := &Log{Dir: dir, Cfg: cfg}
	if _, err := os.Stat(dir); errors.Is(err, fs.ErrNotExist) {
		// The dir does not exist create it
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, err
		}
	}
	return l, l.init()
}

func (l *Log) init() error {
	de, err := os.ReadDir(l.Dir)
	if err != nil {
		return err
	}
	var offsets []uint64
	var osetSet map[uint64]struct{} = make(map[uint64]struct{})

	for _, ent := range de {
		if ent.IsDir() {
			continue // skip over directories
		}
		if storeOff, err := getOffset(ent.Name()); err == nil {
			if _, ok := osetSet[storeOff]; !ok {
				offsets = append(offsets, storeOff)
			}
			osetSet[storeOff] = struct{}{}
		}
	}
	//sort the slice ascending, the newest should come last
	slices.Sort(offsets)
	// We will have 2 offsets one for index and one for store.
	for _, boff := range offsets {
		l.mksegment(boff)
	}
	clear(osetSet)
	if len(l.segments) == 0 {
		err := l.mksegment(l.Cfg.Segment.InitialOffset)
		if err != nil {
			return err
		}
	}
	return nil
}

func getOffset(f string) (uint64, error) {
	offStr := strings.TrimSuffix(filepath.Base(f), filepath.Ext(f))
	off, err := strconv.Atoi(offStr)
	if err != nil {
		return 0, err
	}
	return uint64(off), nil
}

// mksegment makes a segment using state stored in Log
func (l *Log) mksegment(offset uint64) error {
	seg, err := newSegment(l.Dir, offset, l.Cfg)
	if err != nil {
		return err
	}
	l.activeSeg = seg
	l.segments = append(l.segments, seg)
	return nil
}

func (l *Log) Append(record *Record) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	off, err := l.activeSeg.Append(record.Value)
	if err != nil {
		return 0, err
	}
	record.Offset = off

	if l.activeSeg.IsMaxed() {
		err = l.mksegment(off + 1)
	}

	return off, err
}
