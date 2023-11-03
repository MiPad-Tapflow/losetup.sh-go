@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build 

adb push losetup.sh /data/local/tmp

adb shell chmod 777 /data/local/tmp/losetup.sh

adb shell  su -c /data/local/tmp/losetup.sh