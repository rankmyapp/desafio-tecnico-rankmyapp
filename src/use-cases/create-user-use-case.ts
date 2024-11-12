import { randomUUID } from "crypto";
import { DispatchEmail } from "../domain/DispatchEmail";
import { SignUpUser } from "../domain/SignUpUser";
import { User } from "../domain/User";
import { UserRepository } from "../domain/UserRepository";
import { EmailAlreadyTaken } from "../errors/EmailAlreadyTaken";
import { InvalidPassword } from "../errors/InvalidPassword";
import { Encrypter } from "../security/Encrypter";

export class CreateUserUseCase {
  constructor(
    private repository: UserRepository,
    private encrypter: Encrypter,
    private dispatchEmail: DispatchEmail
  ) {}

  async handle(user: SignUpUser): Promise<void> {
    if (user.password.length < 6)
      throw new InvalidPassword("Password must have 6 caracters");
    const userExists = await this.repository.findByEmail(user.email);
    if (userExists) {
      throw new EmailAlreadyTaken(`Email ${user.email} already taken!`);
    }
    const encrypted = await this.encrypter.encrypt(user.password);
    const newUser: User = {
      ...user,
      password: encrypted,
      id: randomUUID(),
    };
    await this.repository.save(newUser);
    await this.dispatchEmail.dispatch(user.email);
  }
}
