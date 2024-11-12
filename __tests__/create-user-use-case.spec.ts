import { describe, expect, it, jest } from "@jest/globals";
import { CreateUserUseCase } from "../src/use-cases/create-user-use-case";
import { User } from "../src/domain/User";
import { InvalidPassword } from "../src/errors/InvalidPassword";
import { UserRepositoryMocks } from "./mocks/UserRepositoryMock";
import { EmailAlreadyTaken } from "../src/errors/EmailAlreadyTaken";
import { EncrypterMock } from "./mocks/EncrypterMock";
import { DispatchEmailMcok } from "./mocks/DispatchEmailMock";
import { SignUpUser } from "../src/domain/SignUpUser";
import { randomUUID } from "crypto";

describe("CreateUserUseCase", () => {
  it("Password has less than 6 caracters", () => {
    const useCase = new CreateUserUseCase(undefined, undefined, undefined);
    const userMock: SignUpUser = {
      email: "email@test",
      password: "123",
      name: "name",
    };
    expect(() => useCase.handle(userMock)).rejects.toThrowError(
      InvalidPassword
    );
  });

  it("Password has less than 6 caracters", () => {
    const userRepoMock = new UserRepositoryMocks();
    const userMock: User = {
      email: "email@test",
      password: "123456",
      name: "name",
      id: randomUUID(),
    };
    jest.spyOn(userRepoMock, "findByEmail").mockResolvedValue(userMock);
    const useCase = new CreateUserUseCase(userRepoMock, undefined, undefined);

    expect(() => useCase.handle(userMock)).rejects.toThrowError(
      EmailAlreadyTaken
    );
  });

  it("Password is encrypted", async () => {
    const userRepoMock = new UserRepositoryMocks();
    const encryptMock = new EncrypterMock();
    const dispatchMock = new DispatchEmailMcok();
    const spyEncrypt = jest.spyOn(encryptMock, "encrypt");
    const userMock: SignUpUser = {
      email: "email@test",
      password: "123456",
      name: "name",
    };
    jest.spyOn(userRepoMock, "findByEmail").mockResolvedValue(null);
    const useCase = new CreateUserUseCase(
      userRepoMock,
      encryptMock,
      dispatchMock
    );
    await useCase.handle(userMock);
    expect(spyEncrypt).toHaveBeenCalled();
  });
});
