# Use the official Golang image as the base image
FROM golang:alpine as build

ENV GIN_MODE=release

# Set the working directory in the container
WORKDIR /go/src/gitgood-backend

# Copy the go.mod and go.sum files into the container
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Build the Go app
RUN go build

# Expose the port that our app will listen on
EXPOSE 8080

# Run the app
ENTRYPOINT ["./gitgood-backend"]
