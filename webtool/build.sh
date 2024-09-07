#!/usr/bin/env bash
# 打包镜像

set -ex

CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o webtool .

docker build -t qvgz/webtool:latest .

docker push qvgz/webtool:latest
