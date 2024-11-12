import { Connection } from "mysql";
import { User } from "../domain/User";
import { UserRepository } from "../domain/UserRepository";

export class UserRepositoryMySQLImpl implements UserRepository {
  constructor(private connection: Connection) {}
  async findByEmail(email: string): Promise<User> {
    return new Promise((resolve, reject) => {
      this.connection.query(
        "SELECT * FROM user WHERE email = ?",
        [email],
        (error, results) => {
          if (error) return reject(error);
          return resolve(results[0]);
        }
      );
    });
  }

  findById(id: string): Promise<User> {
    return new Promise((resolve, reject) => {
      this.connection.query(
        "SELECT * FROM user WHERE id = ?",
        [id],
        (error, results) => {
          if (error) return reject(error);
          return resolve(results[0]);
        }
      );
    });
  }

  save(user: User): Promise<void> {
    return new Promise((resolve, reject) => {
      this.connection.query(
        "INSERT INTO user SET ?",
        user,
        (error) => {
          if (error) return reject(error);
          return resolve();
        }
      );
    });
  }
}
