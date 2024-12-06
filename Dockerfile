FROM golang:1.22.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
RUN go mod tidy 

COPY . .
RUN go build -o main .

FROM ubuntu:latest
WORKDIR /app

COPY --from=builder /app/main main
COPY ./config/conf.toml config/conf.toml
COPY ./simulator.db simulator.db
RUN ls -l

CMD ["main"]