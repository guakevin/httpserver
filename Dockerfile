
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY [ "go.mod", "go.sum", "./" ]
RUN go mod download 

ADD . .

RUN go mod download

RUN go build -o dist/server main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=build /app/dist/server /app/
CMD [ "/app/server", "-p", "8000" ]