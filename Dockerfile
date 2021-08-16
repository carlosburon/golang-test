FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /lana-sre-challenge

EXPOSE 3000
EXPOSE 2112

CMD [ "/lana-sre-challenge"]


