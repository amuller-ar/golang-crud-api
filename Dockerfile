FROM golang:1.16.4-alpine AS builder

RUN apk --no-cache add make git gcc libtool musl-dev dumb-init


# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY . .

RUN go get -u -d -v

# Copy the code into the container
COPY . .

# Build the application
RUN go build -tags netgo -a -v -o main

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM scratch

COPY --from=builder /dist/main /

EXPOSE 8080
# Command to run
ENTRYPOINT ["/main"]