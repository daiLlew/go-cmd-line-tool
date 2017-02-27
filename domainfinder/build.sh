#!/usr/bin/env bash

projects=("domainfinder" "sprinkle" "domainify" "coolify" "synonyms" "thesaurus" "available")

echo Building domainfinder...

for _proj in "${projects[@]}"
do
    echo "    building $_proj..."
    cd "../$_proj"
    go build -o "../domainfinder/lib/$_proj"
done

cd ../domainfinder
cp ../sprinkle/resources/transformations.txt resources/transformations.txt
echo Build Complete.