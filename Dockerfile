
FROM golang:1.18-alpine

WORKDIR /app

ADD . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -installsuffix cgo -o server cmd/api/main.go

EXPOSE 8000

CMD ["./server"]