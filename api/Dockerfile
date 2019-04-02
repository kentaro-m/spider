FROM golang:1.12.1

WORKDIR /go/src/main
COPY . .

ENV GO111MODULE=on

RUN go get -d -v ./...
RUN go get -u github.com/oxequa/realize
RUN go install -v ./...

EXPOSE 8080

CMD ["realize", "start", "--no-config", "--run"]