#!/usr/bin/env bash

for file in cmd/*.go; do
    echo "Building:" ${file##*/}
    name=$(echo "${file##*/}" | cut -f 1 -d '.')
    GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o build/${name} cmd/${file##*/}
    cd build/
    zip ${name##*/}.zip ${name##*/}
    cd -
done