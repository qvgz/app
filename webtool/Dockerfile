FROM golang:alpine AS builder

WORKDIR /tmp/webtool

COPY . .

RUN set -eux \
    ; GO_ENABLED=0 go build -a -trimpath -ldflags "-s -w -buildid= -extldflags -static" -o webtool 


FROM alpine:latest

RUN set -eux ; \
    apk add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime ; \
    echo "Asia/Shanghai" > /etc/timezone ; \
    apk del tzdata

WORKDIR /app

COPY --chmod=0755 --from=builder /tmp/webtool/webtool /app/
COPY README.md  /app/
COPY templates /app/templates

EXPOSE 8081

CMD ["./webtool"]
