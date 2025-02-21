FROM golang:alpine

RUN apk update && apk add \
    curl \
    build-base \
    automake \
    autoconf \
    libtool \
    flex \
    bison \
    unzip

COPY scripts/Makefile /goclips/

RUN cd /goclips && \
    make \
    make install
