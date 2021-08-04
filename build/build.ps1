#!/usr/bin/env pwsh
# Copyright (c) 2021 crosstools. MIT license.
# Please use this only for Windows! Use 'build.sh' for Unix-like operating systems instead!

$Platforms = @{
    windows = "386", "amd64", "arm"
    darwin = "amd64", "arm64"
    linux = "386", "amd64", "arm", "arm64", "mips", "mips64", "mips64le", "mipsle", "ppc64", "ppc64le", "riscv64", "s390x"
}

# gc (golang) exists in this system
if (Get-Command "go.exe" -ErrorAction SilentlyContinue) {
    foreach ($Platform in $Platforms.GetEnumerator()) {
        for ($i = 0; $i -lt $Platform.Value.Count; $i++) {
            $Env:GOOS = $Platform.Name; $Env:GOARCH = $Platform.Value[$i]; go build -o ("crosstools-$($Platform.Name)-$($Platform.Value[$i])")
            Write-Progress -Activity ("Compiling $($Platform.Name)") -Status 'Progress->' -PercentComplete ($i + 1)
        }
    }
    Write-Output "CrossTools was built successfully in the root directory"
} else {
    Write-Error "Error: gc (golang) doesn't exist in this system, please install download Go (https://golang.org/dl/)"
    exit 1
}

