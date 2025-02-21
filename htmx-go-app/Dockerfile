FROM golang:1.23-alpine

WORKDIR /app

EXPOSE 8080

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh
# ENTRYPOINT [ "./entrypoint.sh" ]

CMD ["/bin/sh", "-c", "./entrypoint.sh"]
