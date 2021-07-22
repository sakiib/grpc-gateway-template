FROM golang:latest as builder
WORKDIR /app
COPY grpc .
EXPOSE 8000
ENTRYPOINT ["./grpc"]