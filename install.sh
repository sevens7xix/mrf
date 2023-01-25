#!/bin/bash

set -e

# First go and checkout if the source code is installed on your system
git clone https://github.com/sevens7xix/mrf.git

# Change directory
cd mrf

# Build the application
go build

# Install it
go install

set +e
