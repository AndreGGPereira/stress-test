# Etapa de construção
FROM golang:1.22.4-alpine
WORKDIR /stress-test

# Copie os arquivos go.mod e go.sum
COPY go.mod ./
COPY go.sum ./

# Baixe todas as dependências do módulo Go
RUN go mod tidy
RUN go mod download

# Copie o código-fonte restante

COPY . .

# Construa o executável
RUN go build -o main ./cmd

ENTRYPOINT ["./stress-test"]