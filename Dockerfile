# Use the official Golang image
FROM golang:1.22.3-alpine

# Set the Working Directory in the Container
WORKDIR /myApp/

# Install Bash(For Alpine)
RUN apk update && apk add bash

# Copy the Project into the Container
COPY . .

# Build the app
RUN go build -o server .

# Metadata
LABEL version="0.0.1"
LABEL projectname="ASCII-ART-DOCKERIZE"

# Run the app
CMD ["./server"]