#!/bin/sh

# Setting up Apkg directories
echo "> Creating essential directories..."
mkdir -p /storage/emulated/0/Apkg/apk
mkdir -p /storage/emulated/0/Apkg/gzip
mkdir -p /storage/emulated/0/Apkg/bin
mkdir -p /storage/emulated/0/Apkg/lib


# Setting up variables
echo "> Exporting Variables..."
export HomeDir=/storage/emulated/0
export ApkgHome=/storage/emulated/0/Apkg
export Apk=/storage/emulated/0/Apkg/apk
export Gzip=/storage/emulatrd/0/Apkg/gzip
export Bin=/storage/emulated/0/Apkg/bin
export Lib=/storage/emulated/0/Apkg/lib

