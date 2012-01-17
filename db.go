package levigo

// #cgo LDFLAGS: -lleveldb
// #include <stdlib.h>
// #include "levigo.h"
import "C"

import (
	"unsafe"
)

type DatabaseError string

func (e DatabaseError) Error() string {
	return string(e)
}

type DB struct {
	Ldb *C.leveldb_t
}

type Range struct {
	Start []byte
	Limit []byte
}

func Open(dbname string, o *Options) (*DB, error) {
	var errStr *C.char
	ldbname := C.CString(dbname)
	defer C.free(unsafe.Pointer(ldbname))

	leveldb := C.leveldb_open(o.Opt, ldbname, &errStr)
	if errStr != nil {
		return nil, DatabaseError(C.GoString(errStr))
	}
	return &DB{leveldb}, nil
}

func DestroyDatabase(dbname string, o *Options) error {
	var errStr *C.char
	ldbname := C.CString(dbname)
	defer C.free(unsafe.Pointer(ldbname))

	C.leveldb_destroy_db(o.Opt, ldbname, &errStr)
	if errStr != nil {
		return DatabaseError(C.GoString(errStr))
	}
	return nil
}

func RepairDatabase(dbname string, o *Options) error {
	var errStr *C.char
	ldbname := C.CString(dbname)
	defer C.free(unsafe.Pointer(ldbname))

	C.leveldb_repair_db(o.Opt, ldbname, &errStr)
	if errStr != nil {
		return DatabaseError(C.GoString(errStr))
	}
	return nil
}

func (db *DB) Put(wo *WriteOptions, key, value []byte) error {
	var errStr *C.char
	// leveldb_put, _get, and _delete call memcpy() (by way of Memtable::Add)
	// when called, so we do not need to worry about these []byte being
	// reclaimed by GC.
	kk := (*C.char)(unsafe.Pointer(&key[0]))
	vv := (*C.char)(unsafe.Pointer(&value[0]))
	C.leveldb_put(db.Ldb, wo.Opt,
		kk, C.size_t(len(key)),
		vv, C.size_t(len(value)),
		&errStr)
	if errStr != nil {
		return DatabaseError(C.GoString(errStr))
	}
	return nil
}

func (db *DB) Get(ro *ReadOptions, key []byte) ([]byte, error) {
	var errStr *C.char
	var vallen C.size_t
	value := C.leveldb_get(db.Ldb, ro.Opt,
		(*C.char)(unsafe.Pointer(&key[0])), C.size_t(len(key)),
		&vallen, &errStr)

	if errStr != nil {
		return nil, DatabaseError(C.GoString(errStr))
	}
	if value == nil {
		return nil, nil
	}
	return C.GoBytes(unsafe.Pointer(value), C.int(vallen)), nil
}

func (db *DB) Delete(wo *WriteOptions, key []byte) error {
	var errStr *C.char
	C.leveldb_delete(db.Ldb, wo.Opt,
		(*C.char)(unsafe.Pointer(&key[0])), C.size_t(len(key)),
		&errStr)
	if errStr != nil {
		return DatabaseError(C.GoString(errStr))
	}
	return nil
}

func (db *DB) Write(wo *WriteOptions, w *WriteBatch) error {
	var errStr *C.char
	C.leveldb_write(db.Ldb, wo.Opt, w.Wbatch, &errStr)
	if errStr != nil {
		return DatabaseError(C.GoString(errStr))
	}
	return nil
}

func (db *DB) Iterator(ro *ReadOptions) *Iterator {
	it := C.leveldb_create_iterator(db.Ldb, ro.Opt)
	return &Iterator{Iter: it}
}

func (db *DB) GetApproximateSizes(ranges []Range) []uint64 {
	starts := make([]*C.char, len(ranges))
	limits := make([]*C.char, len(ranges))
	startLens := make([]C.size_t, len(ranges))
	limitLens := make([]C.size_t, len(ranges))
	for i, r := range ranges {
		starts[i] = C.CString(string(r.Start))
		startLens[i] = C.size_t(len(r.Start))
		limits[i] = C.CString(string(r.Limit))
		limitLens[i] = C.size_t(len(r.Limit))
	}
	sizes := make([]uint64, len(ranges))
	numranges := C.int(len(ranges))
	startsPtr := &starts[0]
	limitsPtr := &limits[0]
	startLensPtr := &startLens[0]
	limitLensPtr := &limitLens[0]
	sizesPtr := (*C.uint64_t)(&sizes[0])
	C.leveldb_approximate_sizes(db.Ldb, numranges, startsPtr, startLensPtr, limitsPtr, limitLensPtr, sizesPtr)
	for i, _ := range ranges {
		C.free(unsafe.Pointer(starts[i]))
		C.free(unsafe.Pointer(limits[i]))
	}
	return sizes
}

func (db *DB) PropertyValue(propName string) string {
	cname := C.CString(propName)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(C.leveldb_property_value(db.Ldb, cname))
}

func (db *DB) NewSnapshot() *C.leveldb_snapshot_t {
	return C.leveldb_create_snapshot(db.Ldb)
}

func (db *DB) ReleaseSnapshot(snap *C.leveldb_snapshot_t) {
	C.leveldb_release_snapshot(db.Ldb, snap)
}

func (db *DB) Close() {
	C.leveldb_close(db.Ldb)
}
