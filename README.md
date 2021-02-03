Demo da utilização do broker RabbitMQ intermediando a comunicação entre serviços isolados.

## Pré Requisitos
- [Docker] (https://www.docker.com/)

## Execução
Após baixar o repositório, deve-se acessar a pasta raiz do projeto e executar
```
  docker-compose up
```
que os serviços serão iniciados.

Para consumir, envie um POST para o endpoit `/people` e no corpo passar os atributos: id, nome, email e telefone.
A API, após receber a requisição, enviará essa mensagem para um segundo serviço, que realizará a persistência em um banco postgresql.
 
