FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o apcp-api .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/apcp-api .

EXPOSE 9000

CMD [ "./app" ]

