package syncx

import (
	"errors"
	"sync"
	"time"
)

var ErrClosed = errors.New("batcher closed")

type Sizer func(any) uint

// Batcher provides an API for accumulating items into a batch for processing.
type Batcher[T any] interface {
	// Put adds items to the batcher.
	Put(T) error

	// Get retrieves a batch from the batcher. This call will block until
	// one of the conditions for a "complete" batch is met.
	Get() ([]T, error)

	// Flush forcibly completes the batch currently being built
	Flush() error

	// Close will dispose off the batcher. Any subsequent calls to Put or Flush
	// will return ErrClosed, calls to Get will return an error iff
	// there are no more ready batches.
	Close()

	// IsClosed will determine if the batcher is disposed
	IsClosed() bool
}

type batcher[T any] struct {
	size      Sizer
	batchChan chan []T
	items     []T
	maxTime   time.Duration
	maxItems  uint
	maxBytes  uint
	nBytes    uint
	lock      sync.Mutex
	closed    bool
}

// ready checks for batch readiness, not thread safe
func (b *batcher[T]) ready() bool {
	if b.maxItems != 0 && uint(len(b.items)) >= b.maxItems {
		return true
	}
	if b.maxBytes != 0 && b.nBytes >= b.maxBytes {
		return true
	}
	return false
}

func (b *batcher[T]) flush() {
	b.batchChan <- b.items
	b.items = make([]T, 0, b.maxItems)
	b.nBytes = 0
}

// drain drains the batchChan (batch queue) without blocking
func (b *batcher[T]) drain() {
	for {
		select {
		case <-b.batchChan:
		default:
			return
		}
	}
}

func (b *batcher[T]) Put(item T) error {
	b.lock.Lock()
	if b.closed {
		b.lock.Unlock()
		return ErrClosed
	}

	b.items = append(b.items, item)
	if b.size != nil {
		b.nBytes += b.size(item)
	}

	if b.ready() {
		b.flush()
	}

	b.lock.Unlock()
	return nil
}

func (b *batcher[T]) Get() ([]T, error) {
	var timeout <-chan time.Time
	if b.maxTime > 0 {
		timeout = time.After(b.maxTime)
	}

	select {
	case items, ok := <-b.batchChan:
		// If there's something on the batch channel, we definitely want that.
		if !ok {
			return nil, ErrClosed
		}
		return items, nil
	case <-timeout:
		for {
			if b.lock.TryLock() {
				// We have a lock, try to read from channel first in case
				// something snuck in
				select {
				case items, ok := <-b.batchChan:
					b.lock.Unlock()
					if !ok {
						return nil, ErrClosed
					}
					return items, nil
				default:
				}

				// If that is unsuccessful, nothing was added to the channel,
				// and the temp buffer can't have changed because of the lock,
				// so grab that
				items := b.items
				b.items = make([]T, 0, b.maxItems)
				b.nBytes = 0
				b.lock.Unlock()
				return items, nil
			}
			// If we didn't get a lock, there are two cases:
			// 1) The batch chan is full.
			// 2) A Put or Flush temporarily has the lock.
			// In either case, trying to read something off the batch chan,
			// and going back to trying to get a lock if unsuccessful
			// works.
			select {
			case items, ok := <-b.batchChan:
				if !ok {
					return nil, ErrClosed
				}
				return items, nil
			default:
			}
		}
	}
}

func (b *batcher[T]) Flush() error {
	b.lock.Lock()
	if b.closed {
		b.lock.Unlock()
		return ErrClosed
	}
	b.flush()
	b.lock.Unlock()
	return nil
}

func (b *batcher[T]) Close() {
	for {
		if b.lock.TryLock() {
			if b.closed {
				b.lock.Unlock()
				return
			}
			b.closed = true
			b.items = nil
			b.nBytes = 0
			b.drain()
			close(b.batchChan)
			b.lock.Unlock()
		} else {
			b.drain()
		}
	}
}

func (b *batcher[T]) IsClosed() bool {
	b.lock.Lock()
	c := b.closed
	b.lock.Unlock()
	return c
}

// New creates a new Batcher using the provided arguments.
// Batch readiness can be determined in three ways:
//   - Maximum number of bytes per batch (requires the items provide their size in bytes)
//   - Maximum number of items per batch
//   - Maximum amount of time waiting for a batch (can cause OOM error)
//
// Values of zero for one of these fields indicate they should not be
// taken into account when evaluating the readiness of a batch.
// This provides an ordering guarantee for any given thread such that if a
// thread places two items in the batcher, Get will guarantee the first
// item is returned before the second, whether before the second in the same
// batch, or in an earlier batch.
func NewBatcher[T any](maxTime time.Duration, maxItems, maxBytes, queueLen uint, nbytes Sizer) (Batcher[T], error) {
	if maxBytes > 0 && nbytes == nil {
		return nil, errors.New("batcher: must provide nbytes function")
	}

	return &batcher[T]{
		maxTime:   maxTime,
		maxItems:  maxItems,
		maxBytes:  maxBytes,
		size:      nbytes,
		items:     make([]T, 0, maxItems),
		batchChan: make(chan []T, queueLen),
	}, nil
}
