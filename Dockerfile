# Dockerfile
FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .

ENTRYPOINT /app/build/metrics

RUN make build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go build`
ENTRYPOINT ["/app/build/metrics"]
