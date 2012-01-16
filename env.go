package leveldb

// #cgo LDFLAGS: -lleveldb
// #include "leveldb-cgo.h"
import "C"

func NewDefaultEnv() *C.leveldb_env_t {
	return C.leveldb_create_default_env()
}

func DestroyEnv(env *C.leveldb_env_t) {
	C.leveldb_env_destroy(env)
}
