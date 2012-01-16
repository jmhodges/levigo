package levigo

// #cgo LDFLAGS: -lleveldb
// #include <stdint.h>
// #include "levigo.h"
import "C"

func NewLRUCache(capacity int) *C.leveldb_cache_t {
	return C.leveldb_cache_create_lru(C.size_t(capacity))
}

func DestroyCache(cache *C.leveldb_cache_t) {
	C.leveldb_cache_destroy(cache)
}
