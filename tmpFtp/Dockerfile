FROM golang:alpine AS builder
WORKDIR /tmpftp
COPY . .
RUN set -eux \
  ; GO_ENABLED=0 go build -a -trimpath -ldflags "-s -w -buildid= -extldflags -static" -o tmpftp

FROM scratch
WORKDIR /dir
COPY --chmod=0755 --from=builder /tmpftp/tmpftp /
EXPOSE 9527
ENTRYPOINT ["/tmpftp"]
CMD [ "-dir=/dir" ]
