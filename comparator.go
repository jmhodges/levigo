package levigo

// #cgo LDFLAGS: -lleveldb
// #include "levigo.h"
import "C"

func DestroyComparator(cmp *C.leveldb_comparator_t) {
	C.leveldb_comparator_destroy(cmp)
}
