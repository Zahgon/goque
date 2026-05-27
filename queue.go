package goque

import (
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

// Queue is a standard FIFO (first in, first out) queue.
type Queue struct {
	sync.RWMutex
	DataDir string
	db      *leveldb.DB
	head    uint64
	tail    uint64
	isOpen  bool
}

// OpenQueue opens a queue if one exists at the given directory. If one
// does not already exist, a new queue is created.
func OpenQueue(dataDir string) (*Queue, error) {
	_ = "STUB: not implemented"

	// Create a new Queue.
	return nil, nil
}

// Open database for the queue.

// Check if this Goque type can open the requested data directory.

// Set isOpen and return.

// Enqueue adds an item to the queue.
func (q *Queue) Enqueue(value []byte) (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if queue is closed.
		nil
}

// Create new Item.

// Add it to the queue.

// Increment tail position.

// EnqueueString is a helper function for Enqueue that accepts a
// value as a string rather than a byte slice.
func (q *Queue) EnqueueString(value string) (*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EnqueueObject is a helper function for Enqueue that accepts any
// value type, which is then encoded into a byte slice using
// encoding/gob.
//
// Objects containing pointers with zero values will decode to nil
// when using this function. This is due to how the encoding/gob
// package works. Because of this, you should only use this function
// to encode simple types.
func (q *Queue) EnqueueObject(value interface{}) (*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EnqueueObjectAsJSON is a helper function for Enqueue that accepts
// any value type, which is then encoded into a JSON byte slice using
// encoding/json.
//
// Use this function to handle encoding of complex types.
func (q *Queue) EnqueueObjectAsJSON(value interface{}) (*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Dequeue removes the next item in the queue and returns it.
func (q *Queue) Dequeue() (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if queue is closed.
		nil
}

// Try to get the next item in the queue.

// Remove this item from the queue.

// Increment head position.

// Peek returns the next item in the queue without removing it.
func (q *Queue) Peek() (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if queue is closed.
		nil
}

// PeekByOffset returns the item located at the given offset,
// starting from the head of the queue, without removing it.
func (q *Queue) PeekByOffset(offset uint64) (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if queue is closed.
		nil
}

// PeekByID returns the item with the given ID without removing it.
func (q *Queue) PeekByID(id uint64) (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if queue is closed.
		nil
}

// Update updates an item in the queue without changing its position.
func (q *Queue) Update(id uint64, newValue []byte) (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if queue is closed.
		nil
}

// Check if item exists in queue.

// Create new Item.

// Update this item in the queue.

// UpdateString is a helper function for Update that accepts a value
// as a string rather than a byte slice.
func (q *Queue) UpdateString(id uint64, newValue string) (*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateObject is a helper function for Update that accepts any
// value type, which is then encoded into a byte slice using
// encoding/gob.
//
// Objects containing pointers with zero values will decode to nil
// when using this function. This is due to how the encoding/gob
// package works. Because of this, you should only use this function
// to encode simple types.
func (q *Queue) UpdateObject(id uint64, newValue interface{}) (*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateObjectAsJSON is a helper function for Update that accepts
// any value type, which is then encoded into a JSON byte slice using
// encoding/json.
//
// Use this function to handle encoding of complex types.
func (q *Queue) UpdateObjectAsJSON(id uint64, newValue interface{}) (*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Length returns the total number of items in the queue.
func (q *Queue) Length() uint64 { _ = "STUB: not implemented"; return 0 }

// Close closes the LevelDB database of the queue.
func (q *Queue) Close() error { _ = "STUB: not implemented"; return nil }

// Check if queue is already closed.

// Close the LevelDB database.

// Reset queue head and tail and set
// isOpen to false.

// Drop closes and deletes the LevelDB database of the queue.
func (q *Queue) Drop() error { _ = "STUB: not implemented"; return nil }

// getItemByID returns an item, if found, for the given ID.
func (q *Queue) getItemByID(id uint64) (*Item, error) {
	_ = "STUB: not implemented"
	// Check if empty or out of bounds.
	return nil, nil
}

// Get item from database.

// init initializes the queue data.
func (q *Queue) init() error {
	// Create a new LevelDB Iterator.
	iter := q.db.NewIterator(nil, nil)
	defer iter.Release()

	// Set queue head to the first item.
	if iter.First() {
		q.head = keyToID(iter.Key()) - 1
	}

	// Set queue tail to the last item.
	if iter.Last() {
		q.tail = keyToID(iter.Key())
	}

	return iter.Error()
}
