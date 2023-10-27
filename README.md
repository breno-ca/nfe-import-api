# NFe Import API - GoLang

## Descrição
Este repositório contém uma API em GoLang que permite a importação de Notas Fiscais Eletrônicas (NFe) com o uso do framework Gin-gonic. A autenticação é feita por meio de JSON Web Tokens (JWT). A aplicação é containerizada com Docker e utiliza o banco de dados MySQL para armazenar os dados das NFe importadas.

#### Serviço para teste das rotas online na AWS
ec2-107-21-58-94.compute-1.amazonaws.com:8080

## Configuração
Certifique-se de seguir as etapas abaixo para configurar corretamente o projeto.

### 1. Banco de Dados
O script de criação do banco de dados está presente no projeto, [insira o nome do script aqui]. Certifique-se de executar este script no seu banco de dados MySQL para criar as tabelas necessárias.

### 2. Arquivo de Configuração
O arquivo `env.json` está localizado na raiz do projeto e contém as configurações necessárias. Certifique-se de adicionar as seguintes informações:

- **Porta**: Defina a porta em que a API será executada.
- **Conexão com o Banco de Dados**: Forneça as informações de conexão com o banco de dados, como o host, nome do banco de dados, usuário e senha.
- **Segredo JWT padrão**: O arquivo `env.json` também possui um segredo JWT padrão. Recomendamos que você altere este valor para algo mais seguro antes de implantar em produção.

Aqui está um exemplo de como o arquivo `env.json` pode ser configurado:

```json
{
    "srv_port": "8080",
    "web_ui": true,
    "open_browser": true,
    "db_config": {
        "db_drive": "mysql",
        "db_host": "",
        "db_port": "3306",
        "db_user": "admin",
        "db_pass": "",
        "db_name": "desafio_tecnico_backend"
    },
    "secret": "SNmObxzznpkelAY1QYnV9PhSnvZbHi2xe2q8U89ksNhcDFq1ezGGqosC4ypFaNjSLOeTYFHWkb8JlyA2hglf2j8N"
}

```

### 3. Execução com Docker
Certifique-se de ter o Docker instalado e em execução. Você pode construir e executar o contêiner da seguinte forma:

```bash
docker build -t nfe-import-api .
docker run -p 8080:8080 nfe-import-api
```

A API estará acessível em `http://localhost:8080`.

## Documentação da API
Você pode encontrar a documentação da API e exemplos de solicitações no arquivo [Postman Collection](main/postman-collection.json) incluído neste repositório.

## Contribuições
Contribuições são bem-vindas! Sinta-se à vontade para abrir problemas ou enviar solicitações pull.

## Licença
Este projeto é distribuído sob a licença [MIT](LICENSE).
