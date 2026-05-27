package goque

// goqueType defines the type of Goque data structure used.
type goqueType uint8

// The possible Goque types, used to determine compatibility when
// one stored type is trying to be opened by a different type.
const (
	goqueStack goqueType = iota
	goqueQueue
	goquePriorityQueue
	goquePrefixQueue
)

// checkGoqueType checks if the type of Goque data structure
// trying to be opened is compatible with the opener type.
//
// A file named 'GOQUE' within the data directory used by
// the structure stores the structure type, using the constants
// declared above.
//
// Stacks and Queues are 100% compatible with each other, while
// a PriorityQueue is incompatible with both.
//
// Returns true if types are compatible and false if incompatible.
func checkGoqueType(dataDir string, gt goqueType) (bool, error) {
	_ = "STUB: not implemented"
	// Set the path to 'GOQUE' file.
	return false, nil
}

// Read 'GOQUE' file for this directory.

// Create byte slice of goqueType.

// Get the saved type from the file.

// Convert the file byte to its goqueType.

// Compare the types.
