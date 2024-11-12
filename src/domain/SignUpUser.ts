import { User } from "./User";

export interface SignUpUser extends Omit<User, "id"> {}
