FROM golang:1.10.1-alpine3.7

RUN apk update
RUN apk add git
RUN go get "github.com/armon/go-socks5"
RUN go get "gopkg.in/telegram-bot-api.v4"

RUN mkdir -p /app/

ADD ./src/* /tmp/build/
RUN go build -o /app/socks5 /tmp/build/*.go

CMD ["/app/socks5"]
