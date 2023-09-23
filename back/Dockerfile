FROM golang:latest
LABEL authors="Nifie"

RUN mkdir "/app"
COPY ./ /app/

WORKDIR /app/cmd/
RUN go build -o main

CMD ["/app/cmd/main"]