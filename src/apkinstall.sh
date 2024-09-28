#!/bin/sh

# Check if user is shell
if ["$USER" = "shell"]; then
   echo "! Shell user detected"

  # Setting up Apkg directories
  echo "> Creating essential directories..."
  mkdir -p /storage/emulated/0/Apkg/apk
  mkdir -p /storage/emulated/0/Apkg/gzip
  mkdir -p /data/local/tmp/apkg/bin
  mkdir -p /data/local/tmp/apkg/lib
  
  # Setting up variables
  echo "> Exporting Variables..."
  export Storage=/storage/emulated/0
  export HomeDir=/data/local/tmp
  export ApkgStore=/storage/emulated/0/Apkg
  export ApkgHome=/data/local/tmp/apkg
  export apk=/storage/emulated/0/Apkg/apk
  export gzip=/storage/emulated/0/Apkg/gzip
  export bin=/data/local/tmp/apkg/bin
  export lib=/data/local/tmp/apkg/lib
fi

# Check if user is u0_* (Termux User)
if ["$USER" = "u0_*"]; then
  echo "! Termux user detected"

  # Setting up Apkg directories
  echo "> Creating essential directories..."
  mkdir -p /storage/emulated/0/Apkg/apk
  mkdir -p /storage/emulated/0/Apkg/gzip
  mkdir -p /data/data/com.termux/files/home/Apkg/bin
  mkdir -p /data/data/com.termux/files/home/Apkg/lib
  
  # Setting up variables
  echo "> Exporting Variables..."
  export Storage=/storage/emulated/0
  export HomeDir=/data/data/com.termux/files/home
  export ApkgStore=/storage/emulated/0/Apkg
  export ApkgHome=/data/data/com.termux/files/home
  export apk=/storage/emulated/0/Apkg/apk
  export gzip=/storage/emulated/0/Apkg/gzip
  export bin=/data/data/com.termux/files/home/Apkg/bin
  export lib=/data/data/com.termux/files/home/Apkg/lib
  
  export PATH=/data/data/com.termux/files/home/Apkg/bin
fi