FROM golang:latest AS builder
WORKDIR /stress-test

# Copie os arquivos go.mod e go.sum
COPY go.mod ./


# Baixe todas as dependências do módulo Go
RUN go mod tidy
RUN go mod download

# Copie o código-fonte restante

COPY . .

# Construa o executável
RUN go build -o stress-test ./cmd

ENTRYPOINT ["./stress-test"]