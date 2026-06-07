# Comandos para utilizar o projeto
– De permisão para executar o script
> chmod +x setup.sh
– A primeira vez execute o script somente a primeira vez
> ./setup.sh 
– Depois rode o programa normalmente, na pasta raiz use o comando
> go run cmd/main.go cmd/wire_gen.go

– Dica: Apague os containers com o comando
docker rm -f $(docker ps -a -p)
– Deletar a pasta .docker para o mysql subir zeradinho
– rode o comando abaixo para subir o mysql e o rabitMQ
docker-compose up -d
– execute o mysql com o docker
docker-compose exec mysql bash
– agora que está dentro do banco de dados, crie a tabela orders com 
mysql -uroot -p orders
– insira a senha que é "root"
– rode o comando para rodar o projeto
go run main.go wire_gen.go


– checar container
rode docker ps
– pegue o nome do container da imagem mysql e rode o comando
docker exec -it <TODO: pegar nome do container> bash
– após entrar no bash use comando abaixo para acessar o banco de dados chamado orders
mysql -uroot -p
– insira a senha que é "root" informada em algum canto que não me lembro
– use o comando 
show database
– caso não tenha nenhum database criado então, crie o database
create database orders;

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
