# --- Build Stage ---
FROM golang:1.25-alpine3.22 AS builder

WORKDIR /app

# Download Go modules
RUN apk add --no-cache git

ARG REPO_ACCESS_TOKEN
RUN git config --global url."https://x-access-token:${REPO_ACCESS_TOKEN}@github.com".insteadOf "https://github.com"
ENV GO111MODULE=on
ENV GOPRIVATE=github.com/fastworkco/*

COPY go.mod go.sum ./
RUN go mod download -x

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app-name cmd/main.go

# --- Runtime Stage ---
FROM alpine:3.22
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy built binary from builder stage
COPY --from=builder /app-name /app/app-name
CMD ["/app/app-name"]
