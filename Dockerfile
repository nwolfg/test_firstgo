FROM golang:latest

WORKDIR /go/src/first
COPY ./first.go .

RUN go build first.go

CMD ["./first"]
