#!/usr/bin/env bash
set -e

BINARY="./go-shell"

echo "Compiling..."
go build -gcflags="all=-N -l" -o go-shell ./app/

./go-shell
