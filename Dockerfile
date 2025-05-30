# Build stage for the backend
FROM golang:1.22-alpine AS backend-builder
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/main.go

# Final stage
FROM alpine:latest
WORKDIR /app

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the pre-built frontend files from the host
COPY backend/static ./static

# Copy the backend binary from the backend-builder stage
COPY --from=backend-builder /app/server .

# Copy any necessary configuration files
COPY backend/.env* ./

# Set default API URL for local development
ARG VITE_API_BASE_URL=http://localhost/api/v1
ENV VITE_API_BASE_URL=$VITE_API_BASE_URL

# Expose the port the app runs on
EXPOSE 80

# Command to run the application
CMD ["./server"]
