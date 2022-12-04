FROM golang:1.19-alpine as builder

ADD ./ /ocron
WORKDIR /ocron

RUN go build -o /tmp/ocron ./cmd/ocron

FROM peakcom/s5cmd:v2.0.0 as s5cmd

FROM alpine:3.17

RUN apk update
RUN apk add postgresql15-client

COPY --from=builder /tmp/ocron /usr/bin/ocron
COPY --from=s5cmd /s5cmd /usr/bin/s5cmd

ENTRYPOINT ["ocron", "-config", "/etc/ocron/config.yml"]
