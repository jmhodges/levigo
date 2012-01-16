include $(GOROOT)/src/Make.inc

TARG=levigo
CGOFILES=\
	leveldb.go\
	batch.go\
	comparator.go\
	cache.go\
	db.go\
	env.go\
	iterator.go\
	options.go

CGO_OFILES=\
	c-db.o

include $(GOROOT)/src/Make.pkg
