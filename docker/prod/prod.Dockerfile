FROM golang:latest

RUN apt update && apt install -y \
    clips


WORKDIR /app

ENTRYPOINT ["/bin/sh","/app/scripts/entrypoint.sh"]