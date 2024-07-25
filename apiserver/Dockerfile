FROM golang:1.22-alpine

WORKDIR /go/src/app

COPY . .

RUN go build -o main main.go

EXPOSE 8000

CMD ["./main"]
