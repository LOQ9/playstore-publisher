# Builder stage
FROM golang:1.16 as builder

# Output dir
RUN mkdir -p /build

# Set the Current Working Directory inside the container
WORKDIR /build

# Copy mod file inside the container
# Copy sum file inside the contaner
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy source inside the container
COPY . .

# Compile output
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/go-play-publisher cmd/gpp/main.go

# Thin stage
FROM alpine:3.14

# Set the Current Working Directory inside the container
WORKDIR /app

# Install dependencies
RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/go-play-publisher /app/go-play-publisher

CMD ["/app/go-play-publisher"]
ENTRYPOINT ["/app/go-play-publisher"]
