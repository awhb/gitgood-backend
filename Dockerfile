# Use the official Golang image as the base image
FROM golang:alpine as builder

ENV GIN_MODE=release

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files into the container
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Builds the Go application with CGO disabled and targets a Linux OS.
# The output is a statically linked binary named main.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the builder image
COPY --from=builder /app/main .

# Expose the port that our app will listen on
EXPOSE 8080

# Run the app
ENTRYPOINT ["./main"]
