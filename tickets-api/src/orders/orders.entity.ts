import { Entity, Column, PrimaryGeneratedColumn } from 'typeorm';

@Entity()
export class Order {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  originId: number; // The entity Id that was purchased

  @Column()
  origin: string; // The entity name that was purchased

  @Column()
  status: string; // purchase status

  @Column()
  createdAt: Date;

  @Column()
  updatedAt: Date;
}
