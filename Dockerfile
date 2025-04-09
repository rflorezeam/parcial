# Etapa 1: build
FROM golang:1.22.12-bullseye AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

# Etapa 2: ejecución
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /app/app /app
COPY .env .env

CMD ["/app"]
