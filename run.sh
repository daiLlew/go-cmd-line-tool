#!/usr/bin/env bash

projects=("sprinkle" "domainify" "coolify")

declare -i len="${#projects[@]}"
declare runCMD

for i in "${!projects[@]}"
do
    dir="${projects[$i]}"
    echo "$dir"

    cd "$dir"
    go build -o "$dir"
    cd ..
done
