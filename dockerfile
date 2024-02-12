FROM golang:1.20 AS build
ARG yhttpcmd_version=master
WORKDIR /tmp/build

ENV GO111MODULE=on
ENV GOPROXY https://goproxy.io,direct
ADD . ./
RUN go mod download

RUN GOOS=linux CGO_ENABLED=0 go build -o /go/bin/yhttpcmd

FROM alpine:3.18.0
LABEL yungkei.yhttpcmd.authors="yungkei"
WORKDIR /yhttpcmd
COPY --from=build /go/bin/yhttpcmd /yhttpcmd/
ADD /static /yhttpcmd/static
ADD yhttpcmd.yaml /yhttpcmd/yhttpcmd.yaml
CMD [ "./yhttpcmd","start" ]