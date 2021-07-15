#!/bin/sh
# Copyright (c) 2021 crosstools. MIT license.
# Please use this only for Unix-like operating systems! Use 'build.ps1' for Windows!

windows="386 amd64 arm"
darwin="amd64 arm64"
linux="386 amd64 arm arm64 mips mips64 mips64le mipsle ppc64 ppc64le riscv64 s390x"

runplatform() {
    GOOS=$1 GOARCH=$2 go build -o "crosstools-$1-$2"
}

# gc (golang) exists in this system
if type "go" > /dev/null; then
    for platform in $windows; do
        runplatform "windows" $platform
    done
    for platform in $darwin; do
        runplatform "darwin" $platform
    done
    for platform in $linux; do
        runplatform "linux" $platform
    done
    echo "CrossTools was built successfully into the root directory"
else
    echo "Error: gc (golang) doesn't exist in this system, please install download Go (https://golang.org/dl/)"
    exit 1
fi
