#!/bin/bash

git checkout master
go doc . > godoc.html
git checkout gh-pages

cat head.html godoc.html footer.html > index.html
rm godoc.html
