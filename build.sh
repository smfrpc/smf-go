#!/bin/bash
set -e
set -x
current=$PWD
this_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root=$(git rev-parse --show-toplevel)
export GOPATH=$root
export GOBIN=$root/bin
cd $root/src/smf
go get ./...

additional_pkgs=(
    "github.com/stretchr/testify"
)
for pkg in ${additional_pkgs[@]}; do
    go get $pkg
done

go build ./...
go test ./...
cd $current
