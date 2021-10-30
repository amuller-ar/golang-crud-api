FROM golang:latest AS builder
RUN apt-get update

ENV GIN_MODE=release
ENV GO111MODULE=on
ENV GO_ENVIRONMENT=prodcution

WORKDIR /go/src
COPY go.mod .
RUN go mod download
COPY . .
RUN go build main.go

FROM scratch
COPY --from=builder /go/src .

EXPOSE 8080

ENTRYPOINT ["./main"]