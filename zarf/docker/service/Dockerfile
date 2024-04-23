# Use the official Golang image as a base
FROM golang:1.22 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and build files
COPY go.mod .

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./app/services/node-api

# Start a new stage from Alpine
FROM alpine:3.19

# Set necessary environment variables
ENV INTERNAL_PORT=8080
ENV EXPOSE_PORT=8080

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build /app/main .

# Expose the port the app runs on
EXPOSE $EXPOSE_PORT

EXPOSE 4000

# Command to run the executable
CMD ["./main"]
