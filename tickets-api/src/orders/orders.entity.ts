import { Ticket } from 'src/tickets/tickets.entity';
import { User } from 'src/users/users.entity';

import {
  Entity,
  Column,
  PrimaryGeneratedColumn,
  OneToOne,
  CreateDateColumn,
  UpdateDateColumn,
  ManyToOne,
  JoinColumn,
} from 'typeorm';

@Entity()
export class Order {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  ticketId: number; // The entity Id that was purchased

  @Column()
  status: string; // purchase status

  @Column()
  userId: number; // The user who made the order

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;

  @OneToOne(() => Ticket, (ticket) => ticket.id)
  @JoinColumn()
  ticket: Ticket;

  @ManyToOne(() => User, (user) => user.orders)
  @JoinColumn({ name: 'userId' })
  user: User;
}
