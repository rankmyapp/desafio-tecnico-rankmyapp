import TicketService from "../services/TicketService.js";

class TicketController {
  async catalog(req, res) {
    try {
      const catalog = TicketService.getCatalog();
      return res.json({ success: true, catalog });
    } catch (err) {
      return res.status(500).json({ success: false, message: err.message });
    }
  }

  async buy(req, res) {
    try {
      const { ticketId, payment_type, userId } = req.body;
      if (!ticketId || !payment_type || !userId) {
        return res.status(400).json({ success: false, message: "Campos obrigatórios: ticketId, payment_type, userId" });
      }

      const sale = TicketService.buyTicket({ ticketId, payment_type, userId });
      return res.status(201).json({ success: true, sale });
    } catch (err) {
      return res.status(400).json({ success: false, message: err.message });
    }
  }
}

export default new TicketController();
