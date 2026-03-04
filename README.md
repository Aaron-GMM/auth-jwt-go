# 🔐 Auth JWT Microservice in Go

Um microsserviço de autenticação construído em **Go (Golang)** focado em performance, segurança e boas práticas de engenharia de software. O projeto utiliza **Clean Architecture** para garantir o isolamento das regras de negócio, facilitando a manutenção e a escalabilidade.

## 🚀 Tecnologias e Ferramentas

* **Linguagem:** Go (v1.25)
* **Framework HTTP:** [Gin Web Framework](https://github.com/gin-gonic/gin)
* **Banco de Dados:** PostgreSQL (v15 Alpine via Docker)
* **ORM:** [GORM](https://gorm.io/)
* **Segurança:** Bcrypt (Hash de senhas) e JWT (JSON Web Tokens via `golang-jwt/jwt/v5`)
* **Infraestrutura:** Docker e Docker Compose
* **Migrações:** Atlas (Ariga GORM Provider)

## 🏗️ Arquitetura

O projeto adota uma estrutura baseada em **Clean Architecture** e **Domain-Driven Design (DDD)**, dividida nas seguintes camadas:

* `domain/entity`: Entidades centrais do negócio e mapeamento do banco de dados.
* `repository`: Camada de acesso a dados (isolando o GORM).
* `service`: Onde residem as regras de negócio puras da aplicação.
* `handler`: Controladores de requisições HTTP (Gin) e DTOs (Data Transfer Objects).
* `infra`: Configurações de banco de dados, variáveis de ambiente, logs e utilitários de segurança/JWT.
* `router`: Definição de rotas e injeção de dependências.

## ⚙️ Pré-requisitos

Antes de começar, certifique-se de ter instalado em sua máquina:
* [Go](https://golang.org/doc/install) (Para rodar localmente)
* [Docker](https://docs.docker.com/get-docker/) e [Docker Compose](https://docs.docker.com/compose/install/)

## 🛠️ Como Executar o Projeto

**1. Clone o repositório:**
```bash
git clone https://github.com/Aaron-GMM/auth-jwt-go
cd auth-jwt-go
```
**2. Configure as Variáveis de Ambiente:**
   Crie um arquivo .env na raiz do projeto com as seguintes variáveis:
   Snippet de código
```bash
user=postgres
password=postgres
dbname=authdb
host=db
port=5432
sslmode=disable
DB_URL=host=db user=postgres password=postgres dbname=authdb port=5432 sslmode=disable
JWT_SECRET=SuaChaveSecretaMuitoForteAqui
```
**3. Execute com Docker Compose:**
   A maneira mais fácil de rodar o banco de dados e a aplicação simultaneamente é via Docker:
   Bash
```bash

docker-compose up -d --build
```

**O serviço estará disponível em http://localhost:8080. O healthcheck configurado no docker-compose.yaml garantirá que a aplicação só inicie após o banco de dados estar 100% pronto.**

## 📡 Endpoints da API

Abaixo estão as rotas públicas disponíveis no microsserviço:

### **📝 Registrar Usuário**

Cria um novo usuário, gera o hash da senha e retorna um Token JWT.

   * Rota: POST /api/v1/register

   Body (JSON):
   
    JSON

    {
      "name": "John Doe",
      "email": "john@example.com",
      "password": "strongpassword123"
    }

    Resposta de Sucesso (201 Created):
    JSON

    {
      "message": "success",
      "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }

### **🔑 Login de Usuário**

Valida as credenciais e retorna um Token JWT.

   * Rota: POST /api/v1/login

   Body (JSON):
    
    JSON:
    {
      "email": "john@example.com",
      "password": "strongpassword123"
    }

Resposta de Sucesso (200 OK):
    
    JSON
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
