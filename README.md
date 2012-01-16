# LevelDB via CGO

## Building

You'll need the leveldb shared library installed on your machine. Right now,
that means applying my [latest patch for leveldb's
Makefile](http://code.google.com/p/leveldb/issues/detail?id=27#c10).

Now, suppose you built the leveldb is in /path/to/lib and the headers
directory of leveldb is in /path/to/include. In your clone of leveldb_cgo, you'll run:

  CGO_CFLAGS="-I/path/to/leveldb/include" CGO_LDFLAGS="-L/path/to/lib" make install

and there you go.
