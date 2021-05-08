FROM golang:latest

RUN mkdir -p go/src/app

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get -u github.com/swaggo/swag/cmd/swag \
  && go get -u github.com/cosmtrek/air

EXPOSE 8080 2345

CMD ["air"]