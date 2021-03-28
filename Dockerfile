FROM golang:alpine

WORKDIR /go/src/data-server
COPY . .

RUN go get -d -v ./...
RUN go install -v cmd/main.go

EXPOSE 9090
CMD ["main"]