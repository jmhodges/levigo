package leveldb

// #cgo LDFLAGS: -lleveldb
// #include <stdlib.h>
// #include "leveldb-cgo.h"
import "C"

import (
	"unsafe"
)

type WriteBatch struct {
	wb *C.leveldb_writebatch_t
}

func NewWriteBatch() *WriteBatch {
	wb := C.leveldb_writebatch_create()
	return &WriteBatch{wb}
}

func DestroyWriteBatch(w *WriteBatch) {
	C.leveldb_writebatch_destroy(w.wb)
}

func (w *WriteBatch) Put(key, value []byte) {
	// FIXME: May be too unsafe if C.leveldb_put does not copy the data or
	// places it on another thread.
	C.leveldb_writebatch_put(w.wb,
		(*C.char)(unsafe.Pointer(&key[0])), C.size_t(len(key)),
		(*C.char)(unsafe.Pointer(&value[0])), C.size_t(len(value)))
}

func (w *WriteBatch) Delete(key []byte) {
	C.leveldb_writebatch_delete(w.wb,
		(*C.char)(unsafe.Pointer(&key[0])), C.size_t(len(key)))
}

func (w *WriteBatch) Clear() {
	C.leveldb_writebatch_clear(w.wb)
}
