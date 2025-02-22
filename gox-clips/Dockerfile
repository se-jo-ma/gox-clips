FROM golang:latest

RUN apt update && apt install -y \
    clips

COPY ./scripts/entrypoint.sh /app/scripts/entrypoint.sh
WORKDIR /app

CMD ["/bin/sh","/app/scripts/entrypoint.sh"]