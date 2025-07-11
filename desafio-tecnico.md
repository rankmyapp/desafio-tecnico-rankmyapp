# Teste Técnico - Desenvolvedor Senior Backend (Node.js)

## Objetivo

Este teste tem como objetivo avaliar suas habilidades em desenvolvimento de aplicações descentralizadas, qualidade de software e entrega com testes. 

Você deverá desenvolver um serviço que também servira recursos através de uma API REST, e terá alguma das operações de CRUD como (Create, Read, Update, Delete) sobre um recurso de **usuários**. e um serviço que irá receber uma mensagem via fila e irá enviar enviar um e-mail

1. Fork este repositório (ou crie um novo) e, ao finalizar, envie o link do repositório para a avaliação.
2. Você deverá utilizar **Node.js** puro com **Express.js**, **Fastify.js** ou **Go**, para criação dos serviços
3. O banco de dados pode ser **MongoDB** ou **PostgreSQL** ou **MySQL**.
4. Utilize **JWT** para autenticação e proteção de rotas.
5. Inclua testes unitários para as principais funcionalidades da aplicação.
6. O uso de bibliotecas populares como **Yup** para validação de dados e **Mongoose** (caso use MongoDB) é recomendado, mas não obrigatório.
7. Escreva um Dockerfile
8. Escreva uma stack de serviços com Docker Compose, contemplando os serviços que você irá precisar.
9. Crie conexão com o RabbitMQ, para envio na fila `email-queue` que o serviço de e-mail irá receber e disparar um e-mail

## Requisitos do Projeto

### Funcionalidades do Serviço de E-mail:

1. **Enviar um e-mail de Bem vindo ao Usuário (Sign Up)**
   - Conectar no RabbitMQ
   - Receber um mensagens (consumer) de envio de e-mail na fila `email-queue`
   - Enviar e-mail via SMTP

### Funcionalidades do Serviço de Autenticação:

1. **Registrar Usuário (Sign Up)**
   - Rota: `POST /api/auth/signup`
   - Campos obrigatórios: `nome`, `email`, `senha`
   - Validação:
     - O campo `email` deve ser único.
     - A senha deve ter no mínimo 6 caracteres.

2. **Buscar Usuário Autenticado**
   - Rota: `GET /api/account/me`
   - Deve retornar os dados do usuário especificado pelo `id`.

3. **Autenticar um usuário**
   - Rota: `POST /api/auth/signin`
   - Permite remover um usuário específico.

4. **Envio do e-mail de bem vindo**
   - Conectar no RabbitMQ
   - Enviar uma mensagem (producer) de envio do e-mail na fila `email-queue` para novos usuários

### Regras de Autenticação:

- Apenas usuários autenticados podem acessar as rotas de listagem, busca, atualização e remoção de usuários.
- Utilize JWT para gerar e validar tokens de autenticação.
- Rota para login:
- **Login**
- Rota: `POST /api/signin`
- Parâmetros: `email`, `senha`
- Retorna: Token JWT válido para autenticação.

### Requisitos Extras (Bônus):

- **Paginação**: Implementar paginação na listagem de usuários.
- **Filtros de Pesquisa**: Permitir a busca de usuários por nome ou e-mail através de query params.
- **Log de Erros**: Implementar um sistema básico de log de erros.

---

## Testes Unitários

Implemente testes para as principais funcionalidades da API, utilizando um framework de testes como **Jest** ou **Mocha**:

- Teste as rotas na autenticação, atualização e deleção de conta.
- Teste a autenticação JWT (Opcional).
- Teste a validação de dados (como e-mail único e tamanho da senha).

---

## O que será avaliado:

- Estrutura e organização do código.
- Boas práticas na implementação da API (RESTful, uso correto de status HTTP, etc.).
- Conhecimento de autenticação com JWT.
- Validação e tratamento de erros.
- Implementação de testes unitários.
- Documentação básica da API (ex.: README com instruções de uso).

---

## Entrega:

Ao finalizar o teste, por favor envie um link para o repositório GitHub com o código desenvolvido.


