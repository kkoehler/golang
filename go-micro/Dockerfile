from golang:1.12

ENV GO111MODULE on

RUN apt-get update && apt-get install -y unzip

# Install the Protoc Compiler
RUN cd /go && \
    wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-linux-x86_64.zip && \
    unzip protoc-3.7.1-linux-x86_64.zip

# Install the Protoc Go compiler
RUN go get -u github.com/golang/protobuf/protoc-gen-go

# Install the Go Micro Extension for Protoc Compiler
RUN go get github.com/micro/protoc-gen-micro

# Install all the Go-Micro stuff
RUN go get github.com/micro/go-micro@v1.1.0

WORKDIR /go/app
