
FROM golang:1.22.2-alpine

RUN apk add --no-cache gcc musl-dev libc-dev linux-headers

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1

RUN go build -o main .

CMD ["./main"]