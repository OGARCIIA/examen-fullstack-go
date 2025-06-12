# Imagen base de Go
FROM golang:1.24.4-alpine

# Instalar dependencias necesarias
RUN apk add --no-cache git

# Setear directorio de trabajo
WORKDIR /app

# Copiar go.mod y go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Build (opcional) → para correrlo como binario:
# RUN go build -o main ./cmd/main.go

# O ejecutarlo con "go run"
CMD ["go", "run", "cmd/main.go"]

# Exponer puerto
EXPOSE 8080
