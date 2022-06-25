#!/bin/bash
target="./slides/_slides.md"
rm ${target}
#cat ./slides/000-*.md >> ${target}
#cat ./slides/010-*.md >> ${target}
#cat ./slides/020-*.md >> ${target}
#cat ./slides/030-*.md >> ${target}
#cat ./slides/040-*.md >> ${target}
#cat ./slides/050-*.md >> ${target}
#cat ./slides/060-*.md >> ${target}
#cat ./slides/07*-*.md >> ${target}
cat ./slides/0??-*.md >> ${target}


PORT=5000 marp --html --server ./slides

