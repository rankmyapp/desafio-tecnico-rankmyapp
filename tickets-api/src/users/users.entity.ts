import { Exclude } from 'class-transformer';
import { Order } from '../orders/orders.entity';
import {
  Entity,
  Column,
  PrimaryGeneratedColumn,
  CreateDateColumn,
  AfterInsert,
  AfterRemove,
  AfterUpdate,
  OneToMany,
} from 'typeorm';

@Entity()
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  email: string;

  @Column()
  @Exclude()
  password: string;

  @CreateDateColumn()
  createdAt: Date;

  @CreateDateColumn()
  updatedAt: Date;

  @OneToMany(() => Order, (order) => order.user)
  orders: Order[];

  @AfterInsert()
  logInsert() {
    console.log(`Inserted with id: ${this.id}`);
  }

  @AfterRemove()
  logRemove() {
    console.log(`Removed with id: ${this.id}`);
  }

  @AfterUpdate()
  logUpdate() {
    console.log(`Updated with id: ${this.id}`);
  }
}
