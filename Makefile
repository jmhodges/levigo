include $(GOROOT)/src/Make.inc

TARG=levigo
CGOFILES=\
	batch.go\
	comparator.go\
	cache.go\
	db.go\
	env.go\
	iterator.go\
	options.go\
	conv.go

include $(GOROOT)/src/Make.pkg
