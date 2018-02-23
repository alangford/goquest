FROM golang:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go get github.com/shomali11/slacker

RUN go build -o main .

CMD ["/app/main"]