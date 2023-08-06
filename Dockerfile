FROM golang:1.19-alpine

COPY . /go/src/app

WORKDIR /go/src/app/cmd/port-service

RUN go build -o app main.go

EXPOSE 8080

CMD ["./app"]
