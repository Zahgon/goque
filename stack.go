package goque

import (
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

// Stack is a standard LIFO (last in, first out) stack.
type Stack struct {
	sync.RWMutex
	DataDir string
	db      *leveldb.DB
	head    uint64
	tail    uint64
	isOpen  bool
}

// OpenStack opens a stack if one exists at the given directory. If one
// does not already exist, a new stack is created.
func OpenStack(dataDir string) (*Stack, error) {
	_ = "STUB: not implemented"

	// Create a new Stack.
	return nil, nil
}

// Open database for the stack.

// Check if this Goque type can open the requested data directory.

// Set isOpen and return.

// Push adds an item to the stack.
func (s *Stack) Push(value []byte) (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if stack is closed.
		nil
}

// Create new Item.

// Add it to the stack.

// Increment head position.

// PushString is a helper function for Push that accepts a
// value as a string rather than a byte slice.
func (s *Stack) PushString(value string) (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// PushObject is a helper function for Push that accepts any
		// value type, which is then encoded into a byte slice using
		// encoding/gob.
		//
		// Objects containing pointers with zero values will decode to nil
		// when using this function. This is due to how the encoding/gob
		// package works. Because of this, you should only use this function
		// to encode simple types.
		nil
}

func (s *Stack) PushObject(value interface{}) (*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PushObjectAsJSON is a helper function for Push that accepts any
// value type, which is then encoded into a JSON byte slice using
// encoding/json.
//
// Use this function to handle encoding of complex types.
func (s *Stack) PushObjectAsJSON(value interface{}) (*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pop removes the next item in the stack and returns it.
func (s *Stack) Pop() (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if stack is closed.
		nil
}

// Try to get the next item in the stack.

// Remove this item from the stack.

// Decrement head position.

// Peek returns the next item in the stack without removing it.
func (s *Stack) Peek() (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if stack is closed.
		nil
}

// PeekByOffset returns the item located at the given offset,
// starting from the head of the stack, without removing it.
func (s *Stack) PeekByOffset(offset uint64) (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if stack is closed.
		nil
}

// PeekByID returns the item with the given ID without removing it.
func (s *Stack) PeekByID(id uint64) (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if stack is closed.
		nil
}

// Update updates an item in the stack without changing its position.
func (s *Stack) Update(id uint64, newValue []byte) (*Item, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if stack is closed.
		nil
}

// Check if item exists in stack.

// Create new Item.

// Update this item in the stack.

// UpdateString is a helper function for Update that accepts a value
// as a string rather than a byte slice.
func (s *Stack) UpdateString(id uint64, newValue string) (*Item, error) {
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
func (s *Stack) UpdateObject(id uint64, newValue interface{}) (*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateObjectAsJSON is a helper function for Update that accepts
// any value type, which is then encoded into a JSON byte slice using
// encoding/json.
//
// Use this function to handle encoding of complex types.
func (s *Stack) UpdateObjectAsJSON(id uint64, newValue interface{}) (*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Length returns the total number of items in the stack.
func (s *Stack) Length() uint64 { _ = "STUB: not implemented"; return 0 }

// Close closes the LevelDB database of the stack.
func (s *Stack) Close() error { _ = "STUB: not implemented"; return nil }

// Check if stack is already closed.

// Close the LevelDB database.

// Reset stack head and tail and set
// isOpen to false.

// Drop closes and deletes the LevelDB database of the stack.
func (s *Stack) Drop() error { _ = "STUB: not implemented"; return nil }

// getItemByID returns an item, if found, for the given ID.
func (s *Stack) getItemByID(id uint64) (*Item, error) {
	_ = "STUB: not implemented"
	// Check if empty or out of bounds.
	return nil, nil
}

// Get item from database.

// init initializes the stack data.
func (s *Stack) init() error {
	// Create a new LevelDB Iterator.
	iter := s.db.NewIterator(nil, nil)
	defer iter.Release()

	// Set stack head to the last item.
	if iter.Last() {
		s.head = keyToID(iter.Key())
	}

	// Set stack tail to the first item.
	if iter.First() {
		s.tail = keyToID(iter.Key()) - 1
	}

	return iter.Error()
}
