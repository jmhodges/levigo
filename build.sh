#!/bin/bash
set -o nounset
set -o errexit

git checkout master
# assumes that the repo is checked out in a GOPATH directory as
# github.com/jmhodges/levigo
godoc -html github.com/jmhodges/levigo > godoc.html
git checkout gh-pages

cat head.html godoc.html footer.html > index.html
rm godoc.html
