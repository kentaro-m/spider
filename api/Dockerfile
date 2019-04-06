FROM golang:1.12.1

WORKDIR /go/src/main

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]