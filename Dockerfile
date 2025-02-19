# --------------------------
# Stage 1: Build the binary
# --------------------------
    FROM golang:1.23 AS builder

    # Create and set the working directory inside the container
    WORKDIR /app
    
    # Copy go.mod and go.sum for dependency resolution
    COPY go.mod go.sum ./
    
    # Download all dependencies early for caching benefits
    RUN go mod download
    
    # Copy the entire project into the container
    COPY . .
    
    # Build the server binary in the /app/cmd/server directory
    WORKDIR /app/cmd/server
    # Disable CGO for a fully static binary suitable for Alpine
    RUN CGO_ENABLED=0 go build -o /server .
    
    # --------------------------
    # Stage 2: Create a small runtime image
    # --------------------------
    FROM alpine:3.18
    
    # Install Python 3 and Go for tasks
    RUN apk add --no-cache go python3 

    # (Optional) Set a working directory in the final image
    WORKDIR /
    
    # Copy the compiled binary from the builder stage
    COPY --from=builder /server /server
    
    # Expose the gRPC port
    EXPOSE 50051
    
    # Set the binary as the entrypoint of the container
    ENTRYPOINT ["/server"]
    