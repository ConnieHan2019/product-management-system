FROM golang:1.22 AS builder

LABEL maintainer="ConnieHan2019<18302010059@fudan.edu.cn>"


WORKDIR /app


ENV GOPATH=/app


ENV TZ=Asia/Shanghai


COPY . .


RUN go mod download


RUN go build -o product-management-system ./cmd/main.go


ENTRYPOINT ["./product-management-system"]


WORKDIR /app