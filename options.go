package levigo

// #cgo LDFLAGS: -lleveldb
// #include "levigo.h"
import "C"

// CompressionOpt is a value for Options.SetCompression.
type CompressionOpt int

// Known compression arguments for Options.SetCompression.
const (
	NoCompression     = CompressionOpt(0)
	SnappyCompression = CompressionOpt(1)
)

// Options represent all of the available options when opening a database with
// Open(). Options should be created with NewOptions().
//
// It is usually with to call SetCache() with a cache object. Otherwise, all
// data will be read off disk.
//
// To prevent memory leaks, DestroyOptions() must be called on an Options when
// the program no longer needs it.
//
type Options struct {
	Opt *C.leveldb_options_t
}

// ReadOptions represent all of the available options when reading from a
// database.
//
// To prevent memory leaks, DestroyReadOptions() must called on a ReadOptions
// when the program no longer needs it
type ReadOptions struct {
	Opt *C.leveldb_readoptions_t
}

// WriteOptions represent all of the available options when writeing from a
// database.
//
// To prevent memory leaks, DestroyWriteOptions() must called on a
// WriteOptions when the program no longer needs it
type WriteOptions struct {
	Opt *C.leveldb_writeoptions_t
}

// NewOptions allocates a new Options object.
//
// To prevent memory leaks, the *Options returned must have DestroyOptions()
// called on it when it is no longer needed by the program.
func NewOptions() *Options {
	opt := C.leveldb_options_create()
	return &Options{opt}
}

// NewReadOptions allocates a new ReadOptions object.
//
// To prevent memory leaks, the *ReadOptions returned must have Close() called
// on it when it is no longer needed by the program.
func NewReadOptions() *ReadOptions {
	opt := C.leveldb_readoptions_create()
	return &ReadOptions{opt}
}

// NewWriteOptions allocates a new WriteOptions object.
//
// To prevent memory leaks, the *WriteOptions returned must have Close()
// called on it when it is no longer needed by the program.
func NewWriteOptions() *WriteOptions {
	opt := C.leveldb_writeoptions_create()
	return &WriteOptions{opt}
}

// Close deallocates the Options, freeing its underlying C struct.
func (o *Options) Close() {
	C.leveldb_options_destroy(o.Opt)
}

// SetComparator sets the comparator to be used for all read and write
// operations.
//
// The comparator that created a database must be the same one (technically,
// one with the same name string) that is used to perform read and write
// operations.
//
// The default *C.leveldb_comparator_t is usually sufficient.
func (o *Options) SetComparator(cmp *C.leveldb_comparator_t) {
	C.leveldb_options_set_comparator(o.Opt, cmp)
}

// SetErrorIfExists, if passed true, will cause the opening of a database that
// already exists to throw an error.
func (o *Options) SetErrorIfExists(error_if_exists bool) {
	eie := boolToUchar(error_if_exists)
	C.leveldb_options_set_error_if_exists(o.Opt, eie)
}

// SetCache places a cache object in the database when a database is opened.
//
// This is usually wise to use.
func (o *Options) SetCache(cache *Cache) {
	C.leveldb_options_set_cache(o.Opt, cache.Cache)
}

// SetEnv sets the Env object for the new database handle.
func (o *Options) SetEnv(env *Env) {
	C.leveldb_options_set_env(o.Opt, env.Env)
}

// SetInfoLog sets a *C.leveldb_logger_t object as the informational logger
// for the database.
func (o *Options) SetInfoLog(log *C.leveldb_logger_t) {
	C.leveldb_options_set_info_log(o.Opt, log)
}

// SetWriteBufferSize sets the number of bytes the database will build up in
// memory (backed by an unsorted log on disk) before converting to a sorted
// on-disk file.
func (o *Options) SetWriteBufferSize(s int) {
	C.leveldb_options_set_write_buffer_size(o.Opt, C.size_t(s))
}

// SetParanoidChecks, when called with true, will cause the database to do
// aggressive checking of the data it is processing and will stop early if it
// detects errors.
//
// See the LevelDB C++ documentation docs for details.
func (o *Options) SetParanoidChecks(pc bool) {
	C.leveldb_options_set_paranoid_checks(o.Opt, boolToUchar(pc))
}

// SetMaxOpenFiles sets the number of files than can be used at once by the
// database.
//
// See the LevelDB C++ documentation for details.
func (o *Options) SetMaxOpenFiles(n int) {
	C.leveldb_options_set_max_open_files(o.Opt, C.int(n))
}

// SetBlockSize sets the approximate size of user data packed per block.
//
// See the LevelDB C++ documentation for details.
func (o *Options) SetBlockSize(s int) {
	C.leveldb_options_set_block_size(o.Opt, C.size_t(s))
}

// SetBlockRestartInterval is the number of keys between restarts points for
// delta encoding keys.
//
// Most clients should leave this parameter alone.
func (o *Options) SetBlockRestartInterval(n int) {
	C.leveldb_options_set_block_restart_interval(o.Opt, C.int(n))
}

// SetCompression sets whether to compress blocks using the specified
// compresssion algorithm.
//
// The default value is SnappyCompression and it is fast
// enough that it is unlikely you want to turn it off. The other option is
// NoCompression.
//
// If the LevelDB library was built without Snappy compression enabled, the
// SnappyCompression setting will be ignored.
func (o *Options) SetCompression(t CompressionOpt) {
	C.leveldb_options_set_compression(o.Opt, C.int(t))
}

// SetCreateIfMissing causes Open to create a new database on disk if it does
// not already exist.
func (o *Options) SetCreateIfMissing(b bool) {
	C.leveldb_options_set_create_if_missing(o.Opt, boolToUchar(b))
}

// Close deallocates the ReadOptions, freeing its underlying C struct.
func (ro *ReadOptions) Close() {
	C.leveldb_readoptions_destroy(ro.Opt)
}

// SetVerifyChecksums, when called with true, will cause all data read from
// underlying storage to verified against corresponding checksums.
//
// See the LevelDB C++ documentation for details.
func (ro *ReadOptions) SetVerifyChecksums(b bool) {
	C.leveldb_readoptions_set_verify_checksums(ro.Opt, boolToUchar(b))
}

// SetFillCache, when called with true, will cause all data read from
// underlying storage to be placed in the database cache, if the cache exists.
//
// It is useful to turn this off on ReadOptions that are used for
// *DB.Iterator(), as it will prevent bulk scans from flushing out live user
// data in the cache.
func (ro *ReadOptions) SetFillCache(b bool) {
	C.leveldb_readoptions_set_fill_cache(ro.Opt, boolToUchar(b))
}

// SetSnapshot causes reads to provided as they were when the passed in
// Snapshot was created by *DB.NewSnapshot().
//
// See the LevelDB C++ documentation for details.
func (ro *ReadOptions) SetSnapshot(snap *C.leveldb_snapshot_t) {
	C.leveldb_readoptions_set_snapshot(ro.Opt, snap)
}

// Close deallocates the WriteOptions, freeing its underlying C struct.
func (wo *WriteOptions) Close() {
	C.leveldb_writeoptions_destroy(wo.Opt)
}

// SetSync, when called with true, will cause each write to be flushed from
// the operating system buffer cache before the write is considered complete.
//
// If called with true, this will signficanly slow down writes. If called with
// false, and the machine crashes, some recent writes may be lost. Note that
// if it is just the process that crashes (i.e., the machine does not reboot),
// no writes will be lost even when SetSync is called with false.

// See the LevelDB C++ documentation for details.
func (wo *WriteOptions) SetSync(b bool) {
	C.leveldb_writeoptions_set_sync(wo.Opt, boolToUchar(b))
}
