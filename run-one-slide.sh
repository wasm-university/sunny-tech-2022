#!/bin/bash
target="./slides/_one_slide.md"
rm ${target}
cat ./slides/000-settings.md >> ${target}
cat ${1} >> ${target}

PORT=5000 marp --html --server ./slides

