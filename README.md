# Tickets API - Concert Ticket Sales System

This project is a RESTful API for selling concert tickets, built with NestJS. It allows users to browse available tickets, make purchases.

## Features

- User authentication (signup/signin)
- Ticket catalog with availability information
- Ticket purchase system
- Order processing with message queue
- API versioning

## Description

## Getting Started

### Prerequisites

- Node.js (v14 or higher)
- npm or yarn
- Docker and Docker Compose (for running the database and Redis)

### Installation

1. Clone the repository

```bash
git clone https://github.com/yourusername/desafio-tecnico-rankmyapp.git
cd desafio-tecnico-rankmyapp/tickets-api
```

2. Install dependencies

```bash
npm install
```

3. Set up environment variables

```bash
touch .env
# Create a new file .env and using the .env.example change the variables as it needed
```

4. Start the database and Redis using Docker

```bash
cd ..
docker-compose up -d
```

The API will be available at http://localhost:3000

## Main Endpoints



#### Sign Up

```
POST /api/v1/users/signup
```

Request body:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "id": 1,
  "email": "user@example.com",
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Sign In

```
POST /api/v1/users/signin
```

Request body:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "id": 1,
  "email": "user@example.com",
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```
I have the following endpoints in a collention that can be imported in your Insomnia or Postman. Here is the link to the collection: [tickets-api-collection.json](tickets-api-collection.json)


I also create a user authentication thro each endpoint so to access them you guys will need a Bearer token in each request. `/api/v1/users/signin`

### Tickets

#### Get Ticket Catalog

```
GET /api/v1/tickets/catalog
```

Headers:
```
Authorization: Bearer <access_token>
```

Response:
```json
[
  {
    "id": 1,
    "type": "General Area",
    "availableUnits": 10,
    "price": 95,
    "name": "General Admission",
    "description": "Standing room only",
    "createdAt": "2023-01-01T00:00:00.000Z",
    "updatedAt": "2023-01-01T00:00:00.000Z"
  },
  {
    "id": 2,
    "type": "Grandstand",
    "availableUnits": 5,
    "price": 175,
    "name": "Grandstand Seating",
    "description": "Reserved seating",
    "createdAt": "2023-01-01T00:00:00.000Z",
    "updatedAt": "2023-01-01T00:00:00.000Z"
  },
  // More tickets...
]
```

#### Create Ticket

```
POST /api/v1/tickets
```

Headers:
```
Authorization: Bearer <access_token>
```

Request body:
```json
{
  "type": "General Area",
  "availableUnits": 10,
  "price": 95,
  "name": "General Admission",
  "description": "Standing room only"
}
```

Response:
```json
{
  "id": 1,
  "type": "General Area",
  "availableUnits": 10,
  "price": 95,
  "name": "General Admission",
  "description": "Standing room only",
  "createdAt": "2023-01-01T00:00:00.000Z",
  "updatedAt": "2023-01-01T00:00:00.000Z"
}
```

#### Buy Ticket

```
POST /api/v1/tickets/buy
```

Headers:
```
Authorization: Bearer <access_token>
```

Request body:
```json
{
  "ticketId": 1,
  "paymentType": "CREDIT_CARD",
  "userId": 1
}
```

Response:
```json
{
  "status": "pendingPayment"
}
```

## Architecture

This application follows a modular architecture using NestJS framework:

- **Controllers**: Handle HTTP requests and responses
- **Services**: Contain business logic
- **Entities**: Define database models
- **DTOs**: Define data transfer objects for validation
- **Queue**: Uses BullMQ for asynchronous processing

All sensitive information is stored in environment variables for security.

## Project setup

```bash
$ npm install
```

## Compile and run the project

```bash
# development
$ npm run start

# watch mode
$ npm run start:dev

# production mode
$ npm run start:prod
```

## Run tests

```bash
# unit tests
$ npm run test

# e2e tests
$ npm run test:e2e

# test coverage
$ npm run test:cov
```

## Deployment

When you're ready to deploy your NestJS application to production, there are some key steps you can take to ensure it runs as efficiently as possible. Check out the [deployment documentation](https://docs.nestjs.com/deployment) for more information.

If you are looking for a cloud-based platform to deploy your NestJS application, check out [Mau](https://mau.nestjs.com), our official platform for deploying NestJS applications on AWS. Mau makes deployment straightforward and fast, requiring just a few simple steps:

```bash
$ npm install -g @nestjs/mau
$ mau deploy
```

With Mau, you can deploy your application in just a few clicks, allowing you to focus on building features rather than managing infrastructure.

## Resources

Check out a few resources that may come in handy when working with NestJS:

- Visit the [NestJS Documentation](https://docs.nestjs.com) to learn more about the framework.
- For questions and support, please visit our [Discord channel](https://discord.gg/G7Qnnhy).
- To dive deeper and get more hands-on experience, check out our official video [courses](https://courses.nestjs.com/).
- Deploy your application to AWS with the help of [NestJS Mau](https://mau.nestjs.com) in just a few clicks.
- Visualize your application graph and interact with the NestJS application in real-time using [NestJS Devtools](https://devtools.nestjs.com).
- Need help with your project (part-time to full-time)? Check out our official [enterprise support](https://enterprise.nestjs.com).
- To stay in the loop and get updates, follow us on [X](https://x.com/nestframework) and [LinkedIn](https://linkedin.com/company/nestjs).
- Looking for a job, or have a job to offer? Check out our official [Jobs board](https://jobs.nestjs.com).

## Support

Nest is an MIT-licensed open source project. It can grow thanks to the sponsors and support by the amazing backers. If you'd like to join them, please [read more here](https://docs.nestjs.com/support).

## Stay in touch

- Author - [Kamil Myśliwiec](https://twitter.com/kammysliwiec)
- Website - [https://nestjs.com](https://nestjs.com/)
- Twitter - [@nestframework](https://twitter.com/nestframework)

## License

Nest is [MIT licensed](https://github.com/nestjs/nest/blob/master/LICENSE).
