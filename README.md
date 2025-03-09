# Full Cycle - Go Export - Desafio Go: Stress test

Stress test é uma solução que visa realiar teste de carga com o envio de request simultâneas.

### 📋 Compilação e Execução

1 . Compilar o código-fonte: 
```
go build -o stress-test  cmd/main.go 
```
2 . Executar o teste de carga:

Substitua os valores de --url, --requests e --concurrency de acordo com suas necessidades.

```
./stress-test --url=http://google.com --requests=100 --concurrency=10
```

## ### 📋 Compilação e Execução com Docker

* Build
```sh
docker build -t stresstest .
```  

* Exemplo de execução, testando o host **example.com**, com **30** requisições sendo **duas** concorrentes por vez:
```sh
docker run stresstest --url=http://google.com --requests=30 --concurrency=2
```  




