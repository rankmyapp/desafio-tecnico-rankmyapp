import express from "express";
import routes from "./routes.js";

const app = express();
app.use(express.json());
app.use(routes);

const PORT = process.env.PORT || 3005;

app.listen(PORT, () => {
  console.log(`🚀 API running on http://localhost:${PORT}`);
});
