#!/bin/bash

# 1. Apagar o container da aplicação se existir
docker rm -f desafio-go-3 2>/dev/null || true

# 2. Deletar a pasta .docker (ajuste o caminho se necessário)
rm -rf .docker

# 3. Subir MySQL e RabbitMQ (e a aplicação, se estiver no docker-compose)
docker-compose up -d

# 4. Esperar o MySQL ficar disponível
echo "Aguardando MySQL iniciar..."
until docker-compose exec mysql mysqladmin ping -uroot -proot --silent &> /dev/null ; do
  sleep 2
done
echo "MySQL está pronto!"