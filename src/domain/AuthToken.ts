export interface AuthToken {
  generateToken(id: string): Promise<string>;
}
