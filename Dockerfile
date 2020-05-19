
# Start from golang v1.11 base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Drop it"

RUN mkdir -p /$GOPATH/src/data-post

# Set the Current Working Directory inside the container
WORKDIR /$GOPATH/src/data-post

ENV secureHead="true"
ENV eUser="mcmpadmin"
ENV ePassword="mcmp@Passwd"
ENV elasticClusterIP="cloud-release-ibm-cloud-brokerage-elk-elasticsearch"
ENV dataSize=10000

ENV GO111MODULE=on
#WORKDIR /app

RUN apt-get update && apt-get install -y curl

RUN apt-get update && apt-get install -y apt-utils

RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
RUN echo "y" | apt-get install -y nodejs

RUN npm install -g csvtojson

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
#COPY . /$GOPATH/src/data-post
COPY . .

# Build the Go app
RUN go build -o data-post .

# Expose port 8081 to the outside world
EXPOSE 9095

# Command to run the executable
CMD ["./data-post"]