package levigo

// #cgo LDFLAGS: -lleveldb
// #include "levigo.h"
import "C"

const (
	NoCompression     = 0
	SnappyCompression = 1
)

// Must be freed with DestroyOptions
type Options struct {
	Opt *C.leveldb_options_t
}

type ReadOptions struct {
	Opt *C.leveldb_readoptions_t
}

type WriteOptions struct {
	Opt *C.leveldb_writeoptions_t
}

func NewOptions() *Options {
	opt := C.leveldb_options_create()
	return &Options{opt}
}

func NewReadOptions() *ReadOptions {
	opt := C.leveldb_readoptions_create()
	return &ReadOptions{opt}
}

func NewWriteOptions() *WriteOptions {
	opt := C.leveldb_writeoptions_create()
	return &WriteOptions{opt}
}

func (o *Options) Close() {
	C.leveldb_options_destroy(o.Opt)
}

func (o *Options) SetComparator(cmp *C.leveldb_comparator_t) {
	C.leveldb_options_set_comparator(o.Opt, cmp)
}

func (o *Options) SetErrorIfExists(error_if_exists bool) {
	eie := boolToUchar(error_if_exists)
	C.leveldb_options_set_error_if_exists(o.Opt, eie)
}

func (o *Options) SetCache(cache *Cache) {
	C.leveldb_options_set_cache(o.Opt, cache.Cache)
}

func (o *Options) SetEnv(env *Env) {
	C.leveldb_options_set_env(o.Opt, env.Env)
}

func (o *Options) SetInfoLog(log *C.leveldb_logger_t) {
	C.leveldb_options_set_info_log(o.Opt, log)
}

func (o *Options) SetWriteBufferSize(s int) {
	C.leveldb_options_set_write_buffer_size(o.Opt, C.size_t(s))
}

func (o *Options) SetParanoidChecks(pc bool) {
	C.leveldb_options_set_paranoid_checks(o.Opt, boolToUchar(pc))
}

func (o *Options) SetMaxOpenFiles(n int) {
	C.leveldb_options_set_max_open_files(o.Opt, C.int(n))
}

func (o *Options) SetBlockSize(s int) {
	C.leveldb_options_set_block_size(o.Opt, C.size_t(s))
}

func (o *Options) SetBlockRestartInterval(n int) {
	C.leveldb_options_set_block_restart_interval(o.Opt, C.int(n))
}

func (o *Options) SetCompression(t int) {
	C.leveldb_options_set_compression(o.Opt, C.int(t))
}

func (o *Options) SetCreateIfMissing(b bool) {
	C.leveldb_options_set_create_if_missing(o.Opt, boolToUchar(b))
}

func (ro *ReadOptions) Close() {
	C.leveldb_readoptions_destroy(ro.Opt)
}

func (ro *ReadOptions) SetVerifyChecksums(b bool) {
	C.leveldb_readoptions_set_verify_checksums(ro.Opt, boolToUchar(b))
}

func (ro *ReadOptions) SetFillCache(b bool) {
	C.leveldb_readoptions_set_fill_cache(ro.Opt, boolToUchar(b))
}

func (ro *ReadOptions) SetSnapshot(snap *C.leveldb_snapshot_t) {
	C.leveldb_readoptions_set_snapshot(ro.Opt, snap)
}

func (wo *WriteOptions) Close() {
	C.leveldb_writeoptions_destroy(wo.Opt)
}

func (wo *WriteOptions) SetSync(b bool) {
	C.leveldb_writeoptions_set_sync(wo.Opt, boolToUchar(b))
}
