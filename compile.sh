#!/usr/bin/env bash

projects=("sprinkle" "domainify" "coolify" "synonyms" "thesaurus" "available")

declare -i len="${#projects[@]}"
declare runCMD

for i in "${!projects[@]}"
do
    dir="${projects[$i]}"
    echo "Compiling... $dir"

    cd "$dir"
    go build -o "$dir"
    cd ..
done
