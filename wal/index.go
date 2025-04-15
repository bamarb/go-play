package wal

import (
	"io"
	"os"

	"github.com/tysonmote/gommap"
)

// A record has an offset and a position in the store file.

var (
	offWidth uint64 = 4 // 4 bytes uint32
	posWidth uint64 = 8
	entWidth        = offWidth + posWidth
)

type index struct {
	file *os.File
	mmap gommap.MMap
	size uint64
}

func newIndex(f *os.File, c Config) (*index, error) {
	idx := &index{file: f}
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	sz := fi.Size()
	idx.size = uint64(sz)
	// Truncate expands the file(by filling \0) to MaxIndexBytes or truncates it.
	if err = os.Truncate(f.Name(), int64(c.Segment.MaxIndexBytes)); err != nil {
		return nil, err
	}
	// mmap the file
	if idx.mmap, err = gommap.Map(
		f.Fd(),
		gommap.PROT_READ|gommap.PROT_WRITE,
		gommap.MAP_SHARED); err != nil {
		return nil, err
	}
	return idx, nil
}

// Write appends the given offset and pos to the index
func (ix *index) Write(offset uint32, pos uint64) error {
	if ix.size+entWidth > uint64(len(ix.mmap)) {
		return io.EOF
	}
	enc.PutUint32(ix.mmap[ix.size:ix.size+offWidth], offset)
	enc.PutUint64(ix.mmap[ix.size+offWidth:ix.size+entWidth], pos)
	ix.size += uint64(entWidth)
	return nil
}

// Read reads the offset and position pair (index value) located at in
// the given offset (in) is relative to the segment's base offset
// offsets begin with 0
func (ix *index) Read(in int64) (out uint32, pos uint64, err error) {
	if ix.size == 0 {
		return 0, 0, io.EOF
	}
	// -1 in means read the end of the file
	if in == -1 {
		out = uint32((ix.size / entWidth) - 1)
	} else {
		out = uint32(in)
	}
	pos = uint64(out) * entWidth
	if pos+entWidth > ix.size {
		return 0, 0, io.EOF
	}
	// read from mmap
	out = enc.Uint32(ix.mmap[pos : pos+offWidth])
	pos = enc.Uint64(ix.mmap[pos+offWidth : pos+entWidth])
	return out, pos, nil
}

// Close closes the index flushing the mmap to the disk and cutting it to size
func (ix *index) Close() error {
	if err := ix.mmap.Sync(gommap.MS_SYNC); err != nil {
		return err
	}
	if err := ix.file.Sync(); err != nil {
		return err
	}
	if err := ix.file.Truncate(int64(ix.size)); err != nil {
		return err
	}
	return nil
}

func (ix *index) Name() string {
	return ix.file.Name()
}
