FROM golang:1.14.2

WORKDIR /app

RUN mkdir bin

RUN mkdir src

COPY . /app/src/CrackTheBet/backend

ENV GOPATH=/app

WORKDIR /app/src/CrackTheBet/backend

RUN go get github.com/google/uuid

RUN go get github.com/gorilla/sessions

RUN go get github.com/labstack/echo

RUN go get github.com/labstack/echo-contrib/session

RUN go get github.com/lib/pq

RUN go get golang.org/x/crypto/bcrypt

RUN go install

WORKDIR /app/src/CrackTheBet/backend

RUN go build

CMD ["./backend"]