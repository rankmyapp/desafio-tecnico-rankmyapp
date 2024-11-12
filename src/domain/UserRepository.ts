import { User } from "./User";

export interface UserRepository {
  save(user: User): Promise<void>;

  findByEmail(email: string): Promise<User>;

  findById(id: string): Promise<User>;
}
