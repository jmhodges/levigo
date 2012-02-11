#!/bin/bash
set -o nounset
set -o errexit

git checkout master
godoc -html . > godoc.html
git checkout gh-pages

cat head.html godoc.html footer.html > index.html
rm godoc.html
