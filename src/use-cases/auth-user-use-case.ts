import { AuthToken } from "../domain/AuthToken";
import { SigninUser } from "../domain/SigninUser";
import { UserRepository } from "../domain/UserRepository";
import { InvalidPassword } from "../errors/InvalidPassword";
import { UserDoesNotExists } from "../errors/UserDoesNotExists";
import { Encrypter } from "../security/Encrypter";

export class AuthUserUseCase {
  constructor(
    private repository: UserRepository,
    private encrypter: Encrypter,
    private authToken: AuthToken
  ) {}
  async handle(signIn: SigninUser): Promise<string> {
    const user = await this.repository.findByEmail(signIn.email);
    if (!user)
      throw new UserDoesNotExists(`User ${signIn.email} do not exists!`);
    const passwordsMatch = this.encrypter.isEqual(
      signIn.password,
      user.password
    );
    if (!passwordsMatch) throw new InvalidPassword("Passwords doesn`t match");

    const token = await this.authToken.generateToken(user.id);
    return token;
  }
}
