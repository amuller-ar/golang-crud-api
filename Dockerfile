FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN apk add --no-cache sqlite-libs sqlite-dev
RUN apk add --no-cache build-base
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

RUN ["chmod", "+x", "/go/src/app"]

ENV GIN_MODE=release
EXPOSE 8080

CMD ["golang-crud-api"]
