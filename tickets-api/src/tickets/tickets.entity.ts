import {
  Entity,
  Column,
  PrimaryGeneratedColumn,
  CreateDateColumn,
  UpdateDateColumn,
} from 'typeorm';

export enum TicketsType {
  generalArea = 'General Area',
  grandStand = 'Grand Stand',
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
    enumName: 'tickets_type_enum'
  })
  type: TicketsType;

  @Column()
  availableUnits: number;

  @Column()
  price: number;

  @Column({ nullable: true })
  name?: string;

  @Column({ nullable: true })
  description?: string;

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;
}
