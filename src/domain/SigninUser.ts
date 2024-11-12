import { User } from "./User";

export interface SigninUser extends Pick<User, "email" | "password"> {}
