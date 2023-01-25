#!/bin/bash

set -e

DIR_PATH = "/mrf"

# First go and checkout if the source code is installed on your system
if [ -d "$DIR_PATH" ]; then
    echo "directory already exists" >&2
    exit 1
else
    git clone https://github.com/sevens7xix/mrf.git
fi

# Change directory
cd mrf

# Build the application
go build

# Install it
go install

set +e
