FROM golang:1.20 AS build
ARG yhttpcmd_version=master
WORKDIR /tmp/build

ENV GO111MODULE=on
ENV GOPROXY https://goproxy.io,direct
ADD . ./
RUN go mod download

RUN GOOS=linux CGO_ENABLED=0 go build -o /go/bin/yhttpcmd

FROM alpine:3.18.0
LABEL maintainer="yungkei"
LABEL description="yhttpcmd is a tiny cmd2http(convert command as http service) server, used to execute local commands via http."
WORKDIR /yhttpcmd
COPY --from=build /go/bin/yhttpcmd /yhttpcmd/
ADD yhttpcmd.yaml /yhttpcmd/yhttpcmd.yaml
CMD [ "./yhttpcmd","start" ]