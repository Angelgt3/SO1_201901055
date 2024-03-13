FROM golang:latest

WORKDIR /app

COPY . .

RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/mysql

RUN go get github.com/rs/cors

RUN go build -o app

EXPOSE 8080

CMD ["./app"]
