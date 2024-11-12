import { User } from "../../src/domain/User";
import { UserRepository } from "../../src/domain/UserRepository";

export class UserRepositoryMocks implements UserRepository {
  async findByEmail(email: string): Promise<User> {
    return {
      id: "1",
      email: "email@test",
      name: "name",
      password: "123456",
    };
  }

  async save(user: User): Promise<void> {}

  async findById(id: string): Promise<User> {
    return {
      id: "1",
      email: "email@test",
      name: "name",
      password: "123456",
    };
  }
}
