FROM golang:1.23-alpine AS builder

ADD ./ /ocron
WORKDIR /ocron

RUN go build -o /tmp/ocron ./cmd/ocron

FROM peakcom/s5cmd:v2.2.2 AS s5cmd

FROM alpine:3.20

RUN apk update
RUN apk add postgresql16-client curl

COPY --from=builder /tmp/ocron /usr/bin/ocron
COPY --from=s5cmd /s5cmd /usr/bin/s5cmd

ENTRYPOINT ["ocron", "-config", "/etc/ocron/config.yml"]
