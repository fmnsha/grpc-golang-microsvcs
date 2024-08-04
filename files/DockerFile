FROM golang:1.22  as builder

COPY . /app
WORKDIR /app

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build main.go
#second stage
FROM debian:buster-slim
WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]