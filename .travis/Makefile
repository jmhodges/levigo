all: test

LEVELDB_VERSION ?= v1.18
SNAPPY_VERSION ?= 1.1.1

export CFLAGS ?= -I$(PWD)/root/snappy-$(SNAPPY_VERSION)/include
export CXXFLAGS ?= -I$(PWD)/root/snappy-$(SNAPPY_VERSION)/include
export LDFLAGS ?= -L$(PWD)/root/snappy-$(SNAPPY_VERSION)/lib
export CGO_CFLAGS ?= -I$(PWD)/root/snappy-$(SNAPPY_VERSION)/include -I$(PWD)/root/leveldb/include
export CGO_LDFLAGS ?= -L$(PWD)/root/snappy-$(SNAPPY_VERSION)/lib -L$(PWD)/root/leveldb -lsnappy
export GOPATH ?= $(PWD)/root/go
export LD_LIBRARY_PATH := $(PWD)/root/snappy-$(SNAPPY_VERSION)/lib:$(PWD)/root/leveldb:$(LD_LIBRARY_PATH)

archives/snappy-$(SNAPPY_VERSION).tar.gz: archives
	curl https://snappy.googlecode.com/files/snappy-$(SNAPPY_VERSION).tar.gz > $@

archives:
	mkdir -v archives

levigo: root/snappy-$(SNAPPY_VERSION)/STAMP root/leveldb/STAMP
	cd ../ && go get -d .
	cd ../ && go build .
	cd ../ && go test -test.v=true .

root:
	mkdir -v root

root/leveldb: root
	cd root && git clone https://github.com/google/leveldb.git

root/leveldb/STAMP: root/leveldb root/snappy-$(SNAPPY_VERSION)/STAMP
	cd root/leveldb && git checkout $(LEVELDB_VERSION)
	$(MAKE) -C root/leveldb
	touch $@

root/snappy-$(SNAPPY_VERSION): archives/snappy-$(SNAPPY_VERSION).tar.gz root
	tar xzvf archives/snappy-$(SNAPPY_VERSION).tar.gz -C root

root/snappy-$(SNAPPY_VERSION)/STAMP: root/snappy-$(SNAPPY_VERSION)
	cd root/snappy-$(SNAPPY_VERSION) && ./configure --prefix=$$(pwd)
	$(MAKE) -C root/snappy-$(SNAPPY_VERSION) install
	touch $@

test: levigo

clean:
	-rm -rf archives
	-rm -rf root

.PHONY: levigo test
