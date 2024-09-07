#!/usr/bin/env bash
# 打包镜像

GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -a -ldflags '-extldflags -static -s -w -buildid=' -o tmpFTP-macos .
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags -static -s -w -buildid=' -o tmpFTP-linux .
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags -static -s -w -buildid=' -o tmpFTP-windows .