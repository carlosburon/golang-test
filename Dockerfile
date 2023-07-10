FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /golang-api-test

EXPOSE 3000

CMD [ "/golang-api-test"]


