# Build stage
FROM golang:alpine AS builder

# Establecer variables de entorno para compilación
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GO111MODULE=on

WORKDIR /app

# Copiar módulos primero para aprovechar el caché de Docker
COPY go.mod ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar el binario
RUN go build -o /app/main .

# Runtime stage
FROM alpine:latest

# Instalar ca-certificates para permitir HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copiar binario compilado y assets
COPY --from=builder /app/main .

# Comando de ejecución
ENTRYPOINT ["./main"]