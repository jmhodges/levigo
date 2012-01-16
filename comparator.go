package leveldb

// #cgo LDFLAGS: -lleveldb
// #include "leveldb-cgo.h"
import "C"

func DestroyComparator(cmp *C.leveldb_comparator_t) {
	C.leveldb_comparator_destroy(cmp)
}
