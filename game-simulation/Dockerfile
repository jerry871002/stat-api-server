FROM golang:1.22.5-bullseye as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o game-simulation main.go

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/game-simulation .

EXPOSE 80

CMD ["./game-simulation"]
