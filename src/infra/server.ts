import express from "express";
import { AuthUserUseCase } from "../use-cases/auth-user-use-case";
import { CreateUserUseCase } from "../use-cases/create-user-use-case";
import { connection } from "./mysql";
const app = express();

function setUpApi() {
  app.get("/", (req, res) => {
    res.send();
  });

  app.post("/auth/sign-up", (req, res) => {
    // const useCase = new CreateUserUseCase();
  });

  app.listen(3000);
}

connection.connect(() => setUpApi());
