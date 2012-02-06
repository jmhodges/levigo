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

// DB is a reusabe handle to a LevelDB database on disk, created by Open.
//
// To avoid memory and file descriptor leaks, call *DB.Close() when you are
// through with the handle.
//
// All methods on a DB instance are thread-safe except for the Close()
// method. Calls to any DB method made after Close() will panic.
type DB struct {
	Ldb *C.leveldb_t
}

// Range is a range of keys in the database. GetApproximateSizes calls with it
// begin at the key Start and end right before the key Limit.
type Range struct {
	Start []byte
	Limit []byte
}

// Open opens a database.
//
// Creating a new database is done by calling SetCreateIfMissing(true) on the *Options passed to Open.
//
// It is usually wise to set a Cache object on the *Options with SetCache() to
// keep recently used data from that database in memory.
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

// DestroyDatabase removes a database entirely, removing everything from the
// filesystem.
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

// RepairDatabase attempts to repair a database.
//
// If the database is unrepairable, an error is returned.
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

// Put writes data associated with a key to the database.
//
// If a nil []byte is passed in as value, it will be returned by Get as an
// zero-length slice.
//
// The key and value byte slices may be reused safely. Put takes a copy of
// them before returning.
func (db *DB) Put(wo *WriteOptions, key, value []byte) error {
	var errStr *C.char
	// leveldb_put, _get, and _delete call memcpy() (by way of Memtable::Add)
	// when called, so we do not need to worry about these []byte being
	// reclaimed by GC.
	var k *C.char
	if len(key) != 0 {
		k = (*C.char)(unsafe.Pointer(&key[0]))
	}
	var v *C.char
	if len(value) != 0 {
		v = (*C.char)(unsafe.Pointer(&value[0]))
	}
	C.leveldb_put(db.Ldb, wo.Opt,
		k, C.size_t(len(key)),
		v, C.size_t(len(value)),
		&errStr)
	if errStr != nil {
		return DatabaseError(C.GoString(errStr))
	}
	return nil
}

// Get returns the data associated with the key from the database.
//
// If the key does not exist in the database, a nil []byte is returned. If the
// key does exist, but the data is zero-length in the database, a zero-length
// []byte will be returned.
//
// The key byte slice may be reused safely. Get takes a copy of
// them before returning.
func (db *DB) Get(ro *ReadOptions, key []byte) ([]byte, error) {
	var errStr *C.char
	var vallen C.size_t
	var k *C.char
	if len(key) != 0 {
		k = (*C.char)(unsafe.Pointer(&key[0]))
	}

	value := C.leveldb_get(db.Ldb, ro.Opt,
		k, C.size_t(len(key)),
		&vallen, &errStr)

	if errStr != nil {
		return nil, DatabaseError(C.GoString(errStr))
	}

	if value == nil {
		return nil, nil
	}
	return C.GoBytes(unsafe.Pointer(value), C.int(vallen)), nil
}

// Delete removes the data associated with the key from the database.
//
// The key byte slice may be reused safely. Delete takes a copy of
// them before returning.
func (db *DB) Delete(wo *WriteOptions, key []byte) error {
	var errStr *C.char
	var k *C.char
	if len(key) != 0 {
		k = (*C.char)(unsafe.Pointer(&key[0]))
	}

	C.leveldb_delete(db.Ldb, wo.Opt,
		k, C.size_t(len(key)),
		&errStr)
	if errStr != nil {
		return DatabaseError(C.GoString(errStr))
	}
	return nil
}

// Write atomically writes the *WriteBatch to disk.
func (db *DB) Write(wo *WriteOptions, w *WriteBatch) error {
	var errStr *C.char
	C.leveldb_write(db.Ldb, wo.Opt, w.wbatch, &errStr)
	if errStr != nil {
		return DatabaseError(C.GoString(errStr))
	}
	return nil
}

// NewIterator returns an *Iterator over the the database that uses the
// ReadOptions given.
//
// Often, this is used for large, offline bulk reads while serving live
// traffic. In that case, it may be wise to disable caching so that the data
// processed by the returned *Iterator does not displace the already cached
// data. This can be done by calling SetFillCache(false) on the *ReadOptions
// before passing it here.
//
// Similiarly, *ReadOptions.SetSnapshot() is also useful.
func (db *DB) NewIterator(ro *ReadOptions) *Iterator {
	it := C.leveldb_create_iterator(db.Ldb, ro.Opt)
	return &Iterator{Iter: it}
}

// GetApproximateSizes returns the approximate number of bytes of file system
// space used by one or more key ranges.
//
// The keys counted will begin at Range.Start and end on the key before
// Range.Limit.
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
	C.levigo_leveldb_approximate_sizes(db.Ldb, numranges, startsPtr, startLensPtr, limitsPtr, limitLensPtr, sizesPtr)
	for i, _ := range ranges {
		C.free(unsafe.Pointer(starts[i]))
		C.free(unsafe.Pointer(limits[i]))
	}
	return sizes
}

// PropertyValue returns the value of a database property.
//
// Examples of properties include "leveldb.stats", "leveldb.sstables",
// and "leveldb.num-files-at-level0".
func (db *DB) PropertyValue(propName string) string {
	cname := C.CString(propName)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(C.leveldb_property_value(db.Ldb, cname))
}

// NewSnapshot creates a new snapshot of the database.
//
// The snapshot, when used in a ReadOptions, provides a consistent view of
// state of the database at teh time the snapshot was created.  database and
// returns a handle to it.
func (db *DB) NewSnapshot() *C.leveldb_snapshot_t {
	return C.leveldb_create_snapshot(db.Ldb)
}

// ReleaseSnapshot removes the snapshot from the database's list of snapshots,
// and deallocates it.
func (db *DB) ReleaseSnapshot(snap *C.leveldb_snapshot_t) {
	C.leveldb_release_snapshot(db.Ldb, snap)
}

// Close closes the database, rendering it unusable for I/O, by deallocating
// the underlying handle.
//
// Any attempts to use the DB after Close is called will panic.
//
func (db *DB) Close() {
	C.leveldb_close(db.Ldb)
}
