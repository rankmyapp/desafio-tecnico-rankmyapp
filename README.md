# 🎟️ Ticket API - Desafio Técnico

Este projeto implementa uma **API de Vendas de Tickets** para um show, conforme especificado no teste técnico.  
A aplicação foi desenvolvida em **Node.js** com **Express** e simula um fluxo de vendas de ingressos, incluindo catálogo, compra e publicação em fila.

---

## 🚀 Tecnologias Utilizadas
- [Node.js](https://nodejs.org/) (v22+)
- [Express](https://expressjs.com/)
- [Nodemon](https://nodemon.io/) (para desenvolvimento)

---

## ⚙️ Configuração do Ambiente

### Pré-requisitos
- Node.js v18 ou superior
- NPM ou Yarn

### Clonar o repositório
```bash
git clone https://github.com/seu-usuario/ticket-api.git
cd ticket-api
```

### Instalar dependências
```bash
npm install
```

---

## ▶️ Execução

### Ambiente de Desenvolvimento
```bash
npm run dev
```
O servidor será iniciado em:
```
http://localhost:3000
```

### Produção
```bash
npm start
```

---

## 📌 Endpoints

### 1. Listar Catálogo de Tickets
**GET** `/api/v1/tickets/catalog`

#### Exemplo de Request
```bash
curl http://localhost:3000/api/v1/tickets/catalog
```

#### Exemplo de Response
```json
{
  "success": true,
  "catalog": [
    { "id": "1", "name": "General Area", "price": 95, "stock": 10 },
    { "id": "2", "name": "Grandstand", "price": 175, "stock": 5 },
    { "id": "3", "name": "VIP", "price": 750, "stock": 2 },
    { "id": "4", "name": "Golden Circle", "price": 1250, "stock": 1 }
  ]
}
```

---

### 2. Comprar Ticket
**POST** `/api/v1/tickets/buy`

#### Body (JSON)
```json
{
  "ticketId": "3",
  "payment_type": "CREDIT_CARD",
  "userId": "user123"
}
```

#### Exemplo de Request (curl)
```bash
curl -X POST http://localhost:3000/api/v1/tickets/buy   -H "Content-Type: application/json"   -d '{"ticketId": "3", "payment_type": "CREDIT_CARD", "userId": "user123"}'
```

#### Exemplo de Response
```json
{
  "success": true,
  "sale": {
    "saleId": "1727462738192",
    "ticketId": "3",
    "ticketName": "VIP",
    "userId": "user123",
    "price": 750,
    "createdAt": "2025-08-27T15:25:38.192Z"
  }
}
```

---

## 📦 Estoque Inicial
| Ticket         | Preço (R$) | Quantidade |
|----------------|------------|------------|
| General Area   | 95         | 10         |
| Grandstand     | 175        | 5          |
| VIP            | 750        | 2          |
| Golden Circle  | 1250       | 1          |

---

## 📩 Fila de Mensagens
Cada venda realizada é publicada em uma fila simulada chamada **`validate-purchase`**.  
No código, isso é representado por um array (`purchaseQueue`) que armazena os eventos de validação.

---

## 🛠️ Estrutura do Projeto
```
src/
├── app.js              # Inicialização da aplicação
├── routes.js           # Definição das rotas
├── controllers/        # Controladores HTTP
│   └── TicketController.js
├── services/           # Regras de negócio
│   └── TicketService.js
└── data/               # Dados simulados
    └── tickets.js
```

---

## 🏗️ Build
Como a aplicação é em Node.js puro, não há necessidade de build.  
Para rodar em produção:
```bash
npm install --production
npm start
```

---

## 👨‍💻 Autor
Desenvolvido por *João Vitor Araujo** ✨
