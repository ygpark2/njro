# syntax=docker/dockerfile:1

# =========================================================
# Stage 1: Build binary (Multi-stage build)
# =========================================================
FROM golang:alpine AS builder

# Install build dependencies (CGO support for SQLite/crypto if needed, git, certs)
RUN apk add --no-cache git ca-certificates tzdata build-base

WORKDIR /workspace

# Cache Go modules layer
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build argument: name of the service (e.g. board, post, comment, account, content, emailer, search)
ARG SERVICE
RUN if [ -z "$SERVICE" ]; then echo "ERROR: --build-arg SERVICE=<service_name> is required" && exit 1; fi

# Compile the target service binary with optimizations (-s -w strip debug info)
RUN if [ -d "./service/${SERVICE}" ]; then \
        TARGET_PKG="./service/${SERVICE}"; \
    elif [ -d "./${SERVICE}" ]; then \
        TARGET_PKG="./${SERVICE}"; \
    else \
        echo "ERROR: directory for service '${SERVICE}' not found" && exit 1; \
    fi && \
    CGO_ENABLED=1 GOOS=linux go build \
        -trimpath \
        -ldflags="-s -w" \
        -o /workspace/bin/server \
        ${TARGET_PKG}

# =========================================================
# Stage 2: Minimal Production Runtime
# =========================================================
FROM alpine:3.21 AS runner

# Install essential runtime utilities (certificates for HTTPS/gRPC TLS, timezone data)
RUN apk --no-cache add ca-certificates tzdata

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy the compiled binary from builder stage
COPY --from=builder /workspace/bin/server /app/server

# Set ownership to non-root user
RUN chown -R appuser:appgroup /app

USER appuser

# Health check & entrypoint
ENTRYPOINT ["/app/server"]
