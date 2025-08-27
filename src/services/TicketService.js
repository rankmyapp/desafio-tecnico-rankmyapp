import { tickets, purchaseQueue } from "../data/tickets.js";

class TicketService {
  getCatalog() {
    return tickets.map(t => ({
      id: t.id,
      name: t.name,
      price: t.price,
      stock: t.stock,
    }));
  }

  buyTicket({ ticketId, payment_type, userId }) {
    if (payment_type !== "CREDIT_CARD") {
      throw new Error("Apenas pagamento com cartão de crédito é aceito.");
    }

    const ticket = tickets.find(t => t.id === ticketId);
    if (!ticket) throw new Error("Ticket não encontrado.");
    if (ticket.stock <= 0) throw new Error("Ticket esgotado.");

    ticket.stock -= 1;

    const sale = {
      saleId: Date.now().toString(),
      ticketId: ticket.id,
      ticketName: ticket.name,
      userId,
      price: ticket.price,
      createdAt: new Date(),
    };

    purchaseQueue.push({ event: "validate-purchase", data: sale });

    return sale;
  }
}

export default new TicketService();
