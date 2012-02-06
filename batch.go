package levigo

// #cgo LDFLAGS: -lleveldb
// #include "levigo.h"
import "C"

import (
	"unsafe"
)

type WriteBatch struct {
	wbatch *C.leveldb_writebatch_t
}

func NewWriteBatch() *WriteBatch {
	wb := C.leveldb_writebatch_create()
	return &WriteBatch{wb}
}

func (w *WriteBatch) Close() {
	C.leveldb_writebatch_destroy(w.wbatch)
}

func (w *WriteBatch) Put(key, value []byte) {
	// leveldb_writebatch_put, and _delete call memcpy() (by way of
	// Memtable::Add) when called, so we do not need to worry about these
	// []byte being reclaimed by GC.
	C.leveldb_writebatch_put(w.wbatch,
		(*C.char)(unsafe.Pointer(&key[0])), C.size_t(len(key)),
		(*C.char)(unsafe.Pointer(&value[0])), C.size_t(len(value)))
}

func (w *WriteBatch) Delete(key []byte) {
	C.leveldb_writebatch_delete(w.wbatch,
		(*C.char)(unsafe.Pointer(&key[0])), C.size_t(len(key)))
}

func (w *WriteBatch) Clear() {
	C.leveldb_writebatch_clear(w.wbatch)
}
