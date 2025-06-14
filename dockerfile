FROM golang:1.24.2 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd

# Etapa final com a mesma imagem base
FROM golang:1.24.2

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
