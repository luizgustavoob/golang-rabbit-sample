Demo da utilização do broker RabbitMQ intermediando a comunicação entre serviços isolados.

## Pré Requisitos
* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)
* [Makefile](https://www.gnu.org/software/make/manual/make.html)

## Execução
Após baixar o repositório, deve-se acessar a pasta raiz do projeto e executar
```
  make up
```
que os serviços serão iniciados.

Para consumir, envie um POST para o endpoit `/people` e no corpo preencha os atributos: `id, nome, email e telefone`.
A API, após receber a requisição, enviará essa mensagem para um segundo serviço, que realizará a persistência em um banco PostgreSQL.
