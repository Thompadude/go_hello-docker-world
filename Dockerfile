FROM golang:latest

WORKDIR /app
COPY cmd/main/main.go .

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./app"]