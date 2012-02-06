package levigo

// #cgo LDFLAGS: -lleveldb
// #include "levigo.h"
import "C"

import (
	"unsafe"
)

type IteratorError string

func (e IteratorError) Error() string {
	return string(e)
}

type Iterator struct {
	Iter *C.leveldb_iterator_t
}

func (it *Iterator) Valid() bool {
	return ucharToBool(C.leveldb_iter_valid(it.Iter))
}

func (it *Iterator) Key() []byte {
	var klen C.size_t
	kdata := C.leveldb_iter_key(it.Iter, &klen)
	return C.GoBytes(unsafe.Pointer(kdata), C.int(klen))
}

func (it *Iterator) Value() []byte {
	var vlen C.size_t
	vdata := C.leveldb_iter_value(it.Iter, &vlen)
	return C.GoBytes(unsafe.Pointer(vdata), C.int(vlen))
}

func (it *Iterator) Next() {
	C.leveldb_iter_next(it.Iter)
}

func (it *Iterator) Prev() {
	C.leveldb_iter_prev(it.Iter)
}

func (it *Iterator) SeekToFirst() {
	C.leveldb_iter_seek_to_first(it.Iter)
}

func (it *Iterator) SeekToLast() {
	C.leveldb_iter_seek_to_last(it.Iter)
}

func (it *Iterator) Seek(key []byte) {
	C.leveldb_iter_seek(it.Iter, (*C.char)(unsafe.Pointer(&key[0])), C.size_t(len(key)))
}

func (it *Iterator) GetError() error {
	var errStr *C.char
	C.leveldb_iter_get_error(it.Iter, &errStr)
	if errStr != nil {
		return IteratorError(C.GoString(errStr))
	}
	return nil
}

func (it *Iterator) Close() {
	C.leveldb_iter_destroy(it.Iter)
	it.Iter = nil
}
