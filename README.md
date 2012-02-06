# levigo

levigo is a Go wrapper for LevelDB.

Questions answered at `golang-nuts@googlegroups.com`.

## Building

You'll need the a shared library build of LevelDB installed on your
machine. Right now, that means applying my [latest patch for LevelDB's
Makefile](http://code.google.com/p/leveldb/issues/detail?id=27#c10).

Now, suppose you built the shared LevelDB library in /path/to/lib and the
headers were installed in /path/to/include. In your clone of levigo, you'll
run:

    CGO_CFLAGS="-I/path/to/leveldb/include" CGO_LDFLAGS="-L/path/to/lib" make install

and there you go.

## Caveats

Comparators must be written in C in your own library. And means to be iterate
over WriteBatches must be similarly created.
