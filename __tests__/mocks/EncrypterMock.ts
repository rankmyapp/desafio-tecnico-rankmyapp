import { Encrypter } from "../../src/security/Encrypter";

export class EncrypterMock implements Encrypter {
  async encrypt(info: string): Promise<string> {
    return info;
  }

  async isEqual(decrypted: string, encrypted: string): Promise<boolean> {
    return decrypted === encrypted;
  }
}
