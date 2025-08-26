import {
  Entity,
  Column,
  PrimaryGeneratedColumn,
  CreateDateColumn,
  UpdateDateColumn,
} from 'typeorm';

enum TicketsType {
  generalArea = 'General Area',
  grandStand = 'Grandstand',
  vip = 'VIP',
  goldenCircle = 'Golden Circle',
}

@Entity()
export class Ticket {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({
    type: 'enum',
    enum: TicketsType,
  })
  type: string;

  @Column()
  availableUnits: number;

  @Column()
  price: number;

  @Column()
  name?: string;

  @Column()
  description?: string;

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;
}
