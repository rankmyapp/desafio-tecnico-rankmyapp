export interface Encrypter {
  encrypt(info: string): Promise<string>;

  isEqual(decrypted: string, encrypted: string): Promise<boolean>;
}
