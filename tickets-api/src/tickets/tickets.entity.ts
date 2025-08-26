import {
  Entity,
  Column,
  PrimaryGeneratedColumn,
  CreateDateColumn,
  UpdateDateColumn,
} from 'typeorm';

export enum TicketsType {
  generalArea = 'generalArea',
  grandStand = 'grandStand',
  vip = 'vip',
  goldenCircle = 'goldenCircle',
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

  @Column()
  name?: string;

  @Column()
  description?: string;

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;
}
