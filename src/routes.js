import { Router } from "express";
import TicketController from "./controllers/TicketController.js";

const routes = Router();

routes.get("/api/v1/tickets/catalog", TicketController.catalog);
routes.post("/api/v1/tickets/buy", TicketController.buy);

export default routes;
