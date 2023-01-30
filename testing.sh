#!/bin/bash

if [[ $1 == "test" ]]; then
    find . -name go.mod -execdir go test ./... \;
elif [[ $1 == "benchmark" ]]; then
    find . -name go.mod -execdir go test ./... -bench=. -run=^# \;
else
    go run main.go Daft Punk;
fi
