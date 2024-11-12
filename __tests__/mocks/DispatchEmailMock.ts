import { DispatchEmail } from "../../src/domain/DispatchEmail";

export class DispatchEmailMcok implements DispatchEmail {
  async dispatch(email: string): Promise<void> {}
}
