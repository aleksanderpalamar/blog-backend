
FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GO111MODULE=on

RUN go build -o main .

EXPOSE 3002

CMD ["./main"]