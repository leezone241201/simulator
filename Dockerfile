FROM golang:1.22.1 AS builder

WORKDIR /app

COPY . .

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

RUN go mod tidy && go build -o main .

FROM ubuntu:latest

COPY --from=builder /app/main /app/main
COPY ./config/conf.toml /app/config/conf.toml

CMD ["/app/main"]