
# Desafio Multithreading Go Expert

Este projeto foi desenvolvido como parte do **Desafio Multithreading do curso Go Expert**. O objetivo é buscar informações de endereço de um CEP utilizando **duas APIs distintas** de forma concorrente e retornar o resultado da API que responder mais rápido.

## **Objetivo do Desafio**

- Realizar requisições simultâneas para as seguintes APIs:
  - [BrasilAPI](https://brasilapi.com.br/api/cep/v1/{cep})
  - [ViaCEP](http://viacep.com.br/ws/{cep}/json/)
- Retornar o resultado da API mais rápida.
- Descartar a resposta mais lenta.
- Exibir o resultado no terminal, com os dados do endereço e o nome da API que respondeu.
- Limitar o tempo de resposta em **1 segundo** e exibir um erro de timeout caso nenhuma API responda a tempo.

---

## **Estrutura do Projeto**

```bash
desafio-multithreading/
├── cmd/                    # Ponto de entrada da aplicação
│   └── main.go              # Inicialização do servidor Echo
├── internal/                # Código interno do projeto
│   ├── api/                 # Handlers da API REST
│   │   └── handler.go        # Lógica do endpoint /cep
│   ├── services/            # Lógica de negócio
│   │   └── cep_service.go    # Função fetchAddress para buscar o CEP
│   └── models/              # Definição das structs
│       └── address.go       # Struct Address
├── go.mod                   # Arquivo de dependências do Go
└── README.md                # Documentação do projeto
```

---

## **Como Funciona**

A aplicação possui uma API REST criada com **Echo**. A rota **`GET /cep/{cep}`** permite consultar um endereço utilizando dois provedores de CEP (BrasilAPI e ViaCEP) de forma concorrente.

- Se uma das APIs responde em até 1 segundo, o resultado da API mais rápida é retornado.
- Se nenhuma API responder a tempo, uma mensagem de **timeout** será exibida.

---

## **Como Executar**

### **Pré-requisitos**
- **Go** instalado (versão 1.20 ou superior).

### **Passos:**

1. Clone o repositório:
   ```bash
   git clone https://github.com/WENDELLDELIMA/go-expert-multithreading.git
   cd go-expert-multithreading
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Execute o projeto:
   ```bash
   go run cmd/main.go
   ```

---

## **Testando a API**

### **Exemplo de Requisição com `curl`:**

```bash
curl http://localhost:8080/cep/08583450
```

### **Resposta:**
Se a requisição for bem-sucedida:
```json
{
  "cep": "08583450",
  "state": "SP",
  "city": "Itaquaquecetuba",
  "neighborhood": "Parque Piratininga",
  "street": "Rua Franca Júnior",
  "api_source": "BrasilAPI"
}
```

Se houver timeout:
```json
{
  "error": "Timeout: Nenhuma API respondeu dentro de 1 segundo"
}
```

---

## **Endpoints Disponíveis**

| **Método** | **Rota**       | **Descrição**                |
|------------|----------------|------------------------------|
| `GET`      | `/cep/{cep}`    | Consulta informações do CEP. |

---

## **Detalhes Técnicos**

### **Multithreading com Goroutines e Select:**

O projeto utiliza **goroutines** para fazer requisições simultâneas às APIs e um **select** para:
- Capturar a resposta da API mais rápida.
- Ignorar a resposta mais lenta.
- Implementar um **timeout de 1 segundo**.

### **Função `fetchAddress`:**

A função `fetchAddress` recebe como parâmetros:
- O **CEP**.
- A **URL** da API.
- O nome da API (**BrasilAPI** ou **ViaCEP**).
- O canal `resultChan` para enviar a resposta.

### **Sleep Aleatório:**

A função `fetchAddress` também simula um delay aleatório entre **0.5s** e **1s** para testar concorrência e latência de APIs:
```go
time.Sleep(time.Duration(500+rand.Intn(500)) * time.Millisecond)
```

---

## **Dependências**

O projeto utiliza:
- [Echo](https://github.com/labstack/echo) - Framework Web para API REST.
- [Go Modules](https://golang.org/doc/go1.11#modules) para gerenciamento de dependências.

---

## **Contribuição**

Se desejar contribuir com melhorias:
1. Faça um fork do repositório.
2. Crie uma branch de feature:
   ```bash
   git checkout -b minha-feature
   ```
3. Após fazer suas alterações:
   ```bash
   git add .
   git commit -m "Descrição da feature"
   git push origin minha-feature
   ```
4. Abra um **pull request** no GitHub.

---

## **Autor**

- **Wendell Lima**  
  [GitHub](https://github.com/WENDELLDELIMA)

---

## **Licença**

Este projeto é open-source e está disponível sob a licença MIT.
