FROM alpine:latest

RUN set -eux ; \
    apk add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime ; \
    echo "Asia/Shanghai" > /etc/timezone ; \
    apk del tzdata

WORKDIR /app

COPY README.md webtool /app/
COPY templates /app/templates

EXPOSE 8081

CMD ["./webtool"]
