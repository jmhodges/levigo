# levigo

levigo is a Go wrapper for LevelDB.

See the godoc for API information. There is a web version of the
[API docs](http://jmhodges.github.com/levigo).

Questions answered at `golang-nuts@googlegroups.com`.

## Building

You'll need the shared library build of LevelDB installed on your
machine. Right now, that means applying my [latest patch for LevelDB's
Makefile](http://code.google.com/p/leveldb/issues/detail?id=27#c11).

Now, suppose you built the shared LevelDB library in /path/to/lib and the
headers were installed in /path/to/include. To install levigo remotely, you'll
run:

    CGO_CFLAGS="-I/path/to/leveldb/include" CGO_LDFLAGS="-L/path/to/lib" go get github.com/jmhodges/levigo

and there you go.

## Caveats

Comparators and WriteBatch iterators must be written in C in your own library.
