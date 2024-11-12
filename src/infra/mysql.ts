import mysql from "mysql";

export const connection = mysql.createConnection({
  host: "localhost",
  database: "auth_db",
  user: "auth_user",
  password: "auth_passwd",
});
