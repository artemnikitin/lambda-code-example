#!/usr/bin/env bash

for f in cmd/*.go; do
    echo "Building:" ${f##*/}
    GOOS=linux go build -ldflags "-w -s" -o build/${f##*/} cmd/${f##*/}
    zip build/${f##*/}.zip build/${f##*/}
done