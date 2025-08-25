import { Ticket } from 'src/tickets/tickets.entity';
import {
  Entity,
  Column,
  PrimaryGeneratedColumn,
  OneToOne,
  CreateDateColumn,
  UpdateDateColumn,
} from 'typeorm';

@Entity()
export class Order {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  ticketId: number; // The entity Id that was purchased

  @Column()
  status: string; // purchase status

  // @Column()
  // units: string; // purchase status

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;

  @OneToOne(() => Ticket, (ticket) => ticket.id)
  ticket: Ticket;
}
