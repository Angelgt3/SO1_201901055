FROM golang:1.17 as builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM debian:buster-slim

WORKDIR /app

COPY --from=builder /app/consumer .

CMD ["./main"]
