# Stage 1: Build the Go binary using the host's platform for speed
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder

WORKDIR /app

# Target OS and Architecture arguments provided automatically by Docker Buildx
ARG TARGETOS
ARG TARGETARCH

# Copy module manifests
COPY src/go.mod src/go.sum ./
RUN go mod download

# Copy full source tree
COPY src/ .

# Build statically linked binary using cross-compilation variables
# CGO_ENABLED=0 ensures the binary doesn't depend on host C libraries
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o server ./cmd/server

# Stage 2: Final lightweight runner image (matches target architecture)
FROM alpine:3.19

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/server .

# Default environment flags (can be overridden at runtime)
ENV PORT=8080
ENV MONGO_URI=mongodb://localhost:27017

EXPOSE 8080

CMD ["./server"]