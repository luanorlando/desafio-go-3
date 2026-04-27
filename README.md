# Desenvolvimento do desafio
Desafio usará como base [código fonte](https://github.com/devfullcycle/goexpert/tree/main/20-CleanArch) onde já tem o projeto bem encaminhado conforme exibido em aulas.

# Objetivo principal 
Fazer a listagem para obter `orders` usando os seguintes serviços:  
- **Endpoint Rest** 
- **Endpoint gRPC**
- **Query GraphQL**


# Clean Architecture: Listagem de Orders (REST, gRPC e GraphQL)

**Objetivo** Neste desafio, você deve implementar a funcionalidade de **Listagem de Orders** em sua aplicação Clean Architecture. O objetivo principal é provar o desacoplamento da arquitetura: você criará um único Use Case (ListOrders) e o exporá através de três interfaces de comunicação diferentes simultaneamente.

**Tecnologias e Padrões**

- **Linguagem:** Go (Golang)
- **Arquitetura:** Clean Architecture
- **Comunicação:** REST, gRPC e GraphQL
- **Infraestrutura:** Docker e Docker Compose

**Requisitos Técnicos:**

**Use Case:** Crie o caso de uso de listagem de pedidos (ListOrdersUseCase).

**Interfaces de Entrada:** Disponibilize o acesso a esse Use Case através de:

- **REST:** Endpoint GET /order.
- **gRPC:** Service ListOrders.
- **GraphQL:** Query ListOrders.

**Banco de Dados:**

- Crie as **migrações** necessárias para criar as tabelas do banco de dados.
- O banco deve ser provisionado via Docker.

**Requisitos de Dockerização (Automação Total)** O avaliador não deve executar nenhum comando manual além do Docker Compose up.

1. **Container da Aplicação:** Você deve criar um Dockerfile para a sua aplicação Go.
2. **Orquestração:** O docker-compose.yaml deve subir o banco de dados e o container da aplicação.
3. **Execução Automática:** Ao rodar o comando docker compose up:
- O banco de dados deve subir.
- **As migrações devem ser aplicadas automaticamente.**
- A aplicação deve iniciar e ficar disponível nas portas configuradas.
- *Atenção:* Garanta que a aplicação aguarde o banco estar pronto antes de tentar rodar as migrações ou iniciar (handling de race condition na inicialização).

**Arquivos Auxiliares**

1. **api.http:** Crie um arquivo api.http na raiz contendo as requisições prontas para:
- Criar uma Order (para popular o banco e testar).
- Listar as Orders (para validar o desafio).

**Entregável**

1. **Link do Repositório:** O link para o seu repositório no GitHub.
2. **README:** O arquivo deve conter:
- O comando único de execução (docker compose up).
- As **portas** em que cada serviço (Web, gRPC, GraphQL) está rodando.

**Regras de Entrega:**

**Repositório Exclusivo (Muito Importante):** Este repositório deve conter **APENAS** o código deste desafio.

- **Não entregue** um repositório "monorepo" contendo pastas de outros cursos ou desafios anteriores. Isso bloqueia o processo de correção automática.

**Branch Principal:** Todo o código deve estar na branch main.
