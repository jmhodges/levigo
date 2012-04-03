# levigo

levigo is a Go wrapper for LevelDB.

The API has been godoc'ed and [is available on the
web](http://gopkgdoc.appspot.com/pkg/github.com/jmhodges/levigo).

Questions answered at `golang-nuts@googlegroups.com`.

## Building

You'll need the shared library build of
[LevelDB](http://code.google.com/p/leveldb/) installed on your machine. The
current LevelDB will build it by default.

Now, if you build LevelDB and put the shared library and headers in one of the
standard places for your OS, you'll be able to simply run:

    go get github.com/jmhodges/levigo

But, suppose you put the shared LevelDB library somewhere weird like
/path/to/lib and the headers were installed in /path/to/include. To install
levigo remotely, you'll run:

    CGO_CFLAGS="-I/path/to/leveldb/include" CGO_LDFLAGS="-L/path/to/lib" go get github.com/jmhodges/levigo

and there you go.

Of course, the rules apply locally with `go build` instead of `go get`.

## Caveats

Comparators and WriteBatch iterators must be written in C in your own
library. This seems like a pain in the ass, but remember that you'll have the
LevelDB C API available to your in your client package when you import levigo.
