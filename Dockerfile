# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="trang <trang.phuong.nguyen@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

ENV GO111MODULE=on

# Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -o main .
# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./bash"]