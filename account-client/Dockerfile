# build
FROM golang:1.16-alpine3.12 as builder
WORKDIR /app
COPY . /app
ENV GO111MODULE=on \
    GOOS=linux \
    CGO_ENABLED=0 \
    GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download
ENTRYPOINT ["go", "test", "-v", "./..."] 