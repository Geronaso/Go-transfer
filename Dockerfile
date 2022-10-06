FROM golang:alpine

RUN mkdir /app

ADD main.go /app
ADD database.db /app
ADD go.sum /app
ADD go.mod /app
ADD internal /app/internal
ADD tests /app/tests

WORKDIR /app

RUN go build -o main .

EXPOSE $PORT

CMD ["/app/main"]