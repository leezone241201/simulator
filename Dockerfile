FROM golang:1.22.1 AS builder

WORKDIR /app

COPY . .

ENV GOPROXY=GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

RUN go build -o main .

FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["main"]