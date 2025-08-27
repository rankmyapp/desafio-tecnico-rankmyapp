import { Exclude } from 'class-transformer';
import { Order } from '../orders/orders.entity';
import { Logger } from '@nestjs/common';
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
  private static readonly logger = new Logger('User');

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
    User.logger.log(`User inserted with id: ${this.id}`);
  }

  @AfterRemove()
  logRemove() {
    User.logger.log(`User removed with id: ${this.id}`);
  }

  @AfterUpdate()
  logUpdate() {
    User.logger.log(`User updated with id: ${this.id}`);
  }
}
