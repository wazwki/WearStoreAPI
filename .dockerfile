FROM golang:1.22.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main cmd/main.go

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main /app/main

COPY .env /.env

COPY /db/migrations /db/migrations 

COPY /docs /docs

EXPOSE ${PORT}

CMD ["/app/main"]