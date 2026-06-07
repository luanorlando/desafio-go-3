#!/bin/bash

# 1. Apagar todos os containers
docker rm -f $(docker ps -aq)

# 2. Deletar a pasta .docker (ajuste o caminho se necessário)
rm -rf .docker

# 3. Subir MySQL e RabbitMQ
docker-compose up -d

# 4. Esperar o MySQL ficar disponível
echo "Aguardando MySQL iniciar..."
until docker-compose exec mysql mysqladmin ping -uroot -proot --silent &> /dev/null ; do
  sleep 2
done
echo "MySQL está pronto!"

# 5. Criar a tabela orders no MySQL
docker-compose exec mysql bash -c "
mysql -uroot -proot -e \"
CREATE DATABASE IF NOT EXISTS orders;
USE orders;
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) PRIMARY KEY,
    price FLOAT,
    tax FLOAT,
    final_price FLOAT
);
\"
"