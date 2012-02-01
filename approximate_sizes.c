#include "levigo.h"

void levigo_leveldb_approximate_sizes(
    leveldb_t* db,
    int num_ranges,
    char** range_start_key, const size_t* range_start_key_len,
    char** range_limit_key, const size_t* range_limit_key_len,
    uint64_t* sizes) {
  leveldb_approximate_sizes(db,
                            num_ranges,
                            (const char* const*)range_start_key,
                            range_start_key_len,
                            (const char* const*)range_limit_key,
                            range_limit_key_len,
                            sizes);
}
