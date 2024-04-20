FROM golang:alpine as builder
ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o ./run ./cmd/service

FROM alpine:latest

RUN apk add --no-cache tzdata
ENV TZ=Asia/Yekaterinburg
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app/

COPY --from=builder /app/run .

ENTRYPOINT ["./run", "-config_path",  "/app/configs/config.yml"]