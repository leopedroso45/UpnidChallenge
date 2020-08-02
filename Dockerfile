FROM golang:rc-buster

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test -v ./...

ENTRYPOINT go run .