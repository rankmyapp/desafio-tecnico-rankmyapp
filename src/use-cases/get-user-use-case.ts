import { User } from "../domain/User";
import { UserRepository } from "../domain/UserRepository";
import { UserDoesNotExists } from "../errors/UserDoesNotExists";

export class GetUserUseCase {
  constructor(private repository: UserRepository) {}

  async handle(id: string): Promise<User> {
    const user = await this.repository.findById(id);

    if (!user) throw new UserDoesNotExists("User does not exist");
    return user;
  }
}
