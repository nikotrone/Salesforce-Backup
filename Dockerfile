
# build stage
FROM golang:alpine AS builder
WORKDIR /go/src/GoS2S3
COPY ./src/GoS2S3/ .
RUN apk add --no-cache git
# RUN go-wrapper download   # "go get -d -v ./..."
RUN go get -d -v golang.org/x/net/html
RUN go get -d -v -u github.com/aws/aws-sdk-go
RUN go install -v .
# COPY ./src/GoS2S3/application-config.json /go/bin/

# final stage
FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/* .
COPY ./src/GoS2S3/application-config.json /app/
ENTRYPOINT ./GoS2S3
LABEL Name=go Version=0.0.1
#EXPOSE 
