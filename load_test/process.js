function selectTicket(context, events, done) {
  const tickets = context.vars.catalogResponse;

  try {
    const parsed = typeof tickets === 'string' ? JSON.parse(tickets) : tickets;

    if (Array.isArray(parsed) && parsed.length > 0) {
      const randomTicket = parsed[Math.floor(Math.random() * parsed.length)];
      context.vars.ticketId = randomTicket.id;
    } else {
      console.warn("Catálogo de tickets vazio ou inválido. Usando fallback ID.");
      context.vars.ticketId = "fallback-ticket-id";
    }
  } catch (err) {
    console.error("Erro ao fazer parse da resposta do catálogo:", err);
    context.vars.ticketId = "fallback-ticket-id";
  }

  return done();
}

module.exports = {
  selectTicket
};
