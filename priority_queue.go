package goque

import (
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// prefixSep is the prefix separator for each item key.
var prefixSep []byte = []byte(":")

// order defines the priority ordering of the queue.
type order int

// Defines which priority order to dequeue in.
const (
	ASC  order = iota // Set priority level 0 as most important.
	DESC              // Set priority level 255 as most important.
)

// priorityLevel holds the head and tail position of a priority
// level within the queue.
type priorityLevel struct {
	head uint64
	tail uint64
}

// length returns the total number of items in this priority level.
func (pl *priorityLevel) length() uint64 { _ = "STUB: not implemented"; return 0 }

// PriorityQueue is a standard FIFO (first in, first out) queue with
// priority levels.
type PriorityQueue struct {
	sync.RWMutex
	DataDir  string
	db       *leveldb.DB
	order    order
	levels   [256]*priorityLevel
	curLevel uint8
	isOpen   bool
}

// OpenPriorityQueue opens a priority queue if one exists at the given
// directory. If one does not already exist, a new priority queue is
// created.
func OpenPriorityQueue(dataDir string, order order) (*PriorityQueue, error) {
	_ = "STUB: not implemented"

	// Create a new PriorityQueue.
	return nil, nil
}

// Open database for the priority queue.

// Check if this Goque type can open the requested data directory.

// Set isOpen and return.

// Enqueue adds an item to the priority queue.
func (pq *PriorityQueue) Enqueue(priority uint8, value []byte) (*PriorityItem, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if queue is closed.
		nil
}

// Get the priorityLevel.

// Create new PriorityItem.

// Add it to the priority queue.

// Increment tail position.

// If this priority level is more important than the curLevel.

// EnqueueString is a helper function for Enqueue that accepts a
// value as a string rather than a byte slice.
func (pq *PriorityQueue) EnqueueString(priority uint8, value string) (*PriorityItem, error) {
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
func (pq *PriorityQueue) EnqueueObject(priority uint8, value interface{}) (*PriorityItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EnqueueObjectAsJSON is a helper function for Enqueue that accepts
// any value type, which is then encoded into a JSON byte slice using
// encoding/json.
//
// Use this function to handle encoding of complex types.
func (pq *PriorityQueue) EnqueueObjectAsJSON(priority uint8, value interface{}) (*PriorityItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Dequeue removes the next item in the priority queue and returns it.
func (pq *PriorityQueue) Dequeue() (*PriorityItem, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if queue is closed.
		nil
}

// Try to get the next item.

// Remove this item from the priority queue.

// Increment head position.

// DequeueByPriority removes the next item in the given priority level
// and returns it.
func (pq *PriorityQueue) DequeueByPriority(priority uint8) (*PriorityItem, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if queue is closed.
		nil
}

// Try to get the next item in the given priority level.

// Remove this item from the priority queue.

// Increment head position.

// Peek returns the next item in the priority queue without removing it.
func (pq *PriorityQueue) Peek() (*PriorityItem, error) { _ = "STUB: not implemented"; return nil, nil }

// Check if queue is closed.

// PeekByOffset returns the item located at the given offset,
// starting from the head of the queue, without removing it.
func (pq *PriorityQueue) PeekByOffset(offset uint64) (*PriorityItem, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Check if queue is closed.
}

// Check if queue is empty.

// If the offset is within the current priority level.

// PeekByPriorityID returns the item with the given ID and priority without
// removing it.
func (pq *PriorityQueue) PeekByPriorityID(priority uint8, id uint64) (*PriorityItem, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Check if queue is closed.
}

// Update updates an item in the priority queue without changing its
// position.
func (pq *PriorityQueue) Update(priority uint8, id uint64, newValue []byte) (*PriorityItem, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if queue is closed.
		nil
}

// Check if item exists in queue.

// Create new PriorityItem.

// Update this item in the queue.

// UpdateString is a helper function for Update that accepts a value
// as a string rather than a byte slice.
func (pq *PriorityQueue) UpdateString(priority uint8, id uint64, newValue string) (*PriorityItem, error) {
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
func (pq *PriorityQueue) UpdateObject(priority uint8, id uint64, newValue interface{}) (*PriorityItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateObjectAsJSON is a helper function for Update that accepts
// any value type, which is then encoded into a JSON byte slice using
// encoding/json.
//
// Use this function to handle encoding of complex types.
func (pq *PriorityQueue) UpdateObjectAsJSON(priority uint8, id uint64, newValue interface{}) (*PriorityItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Length returns the total number of items in the priority queue.
func (pq *PriorityQueue) Length() uint64 { _ = "STUB: not implemented"; return 0 }

// Close closes the LevelDB database of the priority queue.
func (pq *PriorityQueue) Close() error { _ = "STUB: not implemented"; return nil }

// Check if queue is already closed.

// Close the LevelDB database.

// Reset head and tail of each priority level
// and set isOpen to false.

// Drop closes and deletes the LevelDB database of the priority queue.
func (pq *PriorityQueue) Drop() error { _ = "STUB: not implemented"; return nil }

// cmpAsc returns wehther the given priority level is higher than the
// current priority level based on ascending order.
func (pq *PriorityQueue) cmpAsc(priority uint8) bool { _ = "STUB: not implemented"; return false }

// cmpAsc returns wehther the given priority level is higher than the
// current priority level based on descending order.
func (pq *PriorityQueue) cmpDesc(priority uint8) bool { _ = "STUB: not implemented"; return false }

// resetCurrentLevel resets the current priority level of the queue
// so the highest level can be found.
func (pq *PriorityQueue) resetCurrentLevel() { _ = "STUB: not implemented"; return }

// findOffset finds the given offset from the current queue position
// based on priority order.
func (pq *PriorityQueue) findOffset(offset uint64) (*PriorityItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle newLevel initialization for descending order.

// For condition expression.

// For loop expression.

// Level comparison.

// Loop through the priority levels.

// If this level is lower than the current level based on ordering and contains items.

// If the offset is within the current priority level.

// getNextItem returns the next item in the priority queue, updating
// the current priority level of the queue if necessary.
func (pq *PriorityQueue) getNextItem() (*PriorityItem, error) {
	_ = "STUB: not implemented"
	// If the current priority level is empty.
	return nil, nil
}

// Set starting value for curLevel.

// Try to get the next priority level.

// If still empty, return queue empty error.

// Try to get the next item in the current priority level.

// getItemByID returns an item, if found, for the given ID.
func (pq *PriorityQueue) getItemByPriorityID(priority uint8, id uint64) (*PriorityItem, error) {
	_ = "STUB: not implemented"
	// Check if empty or out of bounds.
	return nil, nil
}

// Get item from database.

// generatePrefix creates the key prefix for the given priority level.
func (pq *PriorityQueue) generatePrefix(level uint8) []byte {
	_ = "STUB: not implemented"
	// priority + prefixSep = 1 + 1 = 2
	return nil
}

// generateKey create a key to be used with LevelDB.
func (pq *PriorityQueue) generateKey(priority uint8, id uint64) []byte {
	_ = "STUB: not implemented"
	// prefix + key = 2 + 8 = 10
	return nil
}

// init initializes the priority queue data.
func (pq *PriorityQueue) init() error {
	// Set starting value for curLevel.
	pq.resetCurrentLevel()

	// Loop through each priority level.
	for i := 0; i <= 255; i++ {
		// Create a new LevelDB Iterator for this priority level.
		prefix := pq.generatePrefix(uint8(i))
		iter := pq.db.NewIterator(util.BytesPrefix(prefix), nil)

		// Create a new priorityLevel.
		pl := &priorityLevel{
			head: 0,
			tail: 0,
		}

		// Set priority level head to the first item.
		if iter.First() {
			pl.head = keyToID(iter.Key()[2:]) - 1

			// Since this priority level has item(s), handle updating curLevel.
			if pq.cmpAsc(uint8(i)) || pq.cmpDesc(uint8(i)) {
				pq.curLevel = uint8(i)
			}
		}

		// Set priority level tail to the last item.
		if iter.Last() {
			pl.tail = keyToID(iter.Key()[2:])
		}

		if iter.Error() != nil {
			return iter.Error()
		}

		pq.levels[i] = pl
		iter.Release()
	}

	return nil
}
