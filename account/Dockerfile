FROM golang:1.16-alpine AS builder

WORKDIR /app

# Set go module support on
ENV GO111MODULE=on

# Copy go mod and sum files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

RUN go get github.com/cespare/reflex

COPY . .

# Build the application
# CGO_ENABLED must be disabled for alpine
ENV CGO_ENABLED=0
RUN go build -o main .

# Final stage
FROM alpine:3.14

WORKDIR /app

# Copy built binary from builder
COPY --from=builder /app/main .

# Copy .env file
# Note that this is not a recommended practice for production
# You should use a secrets management tool
# or environment variables passed from your orchestrator
COPY .env.dev .

# Expose port
EXPOSE 8080

# We use reflex to watch for file changes in development
# For production, this should be changed to "./main"
# CMD [ "./main" ]
CMD ["reflex", "-c", "reflex.conf"]