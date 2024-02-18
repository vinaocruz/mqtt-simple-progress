# Simple Progress com protocolo MQTT

Este projeto é uma simples prova de conceito sobre uso do RabbitMQ como mqtt broker para implementação do pub/sub com go e javascript, respectivamente.

## Setup

Execute o docker-compose para subir o broker com rabbitmq

```
docker-compose up -d
```

## Publisher

O publisher está implementado com go na pasta publisher. Abra no terminal e digite um valor numérico de progresso (ex: valores entre 0 e 100%)

```

```

## Subscriber

Usando a biblioteca [MQTT.js](https://github.com/mqttjs), o subscriver estará aguardando a mensagem publicada.

```

```