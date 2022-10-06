FROM golang:1.19

RUN mkdir /app

ADD main.go /app
ADD database.db /app
ADD go.sum /app
ADD go.mod /app
ADD internal /app/internal

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]