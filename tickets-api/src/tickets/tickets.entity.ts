import { Entity, Column, PrimaryGeneratedColumn } from 'typeorm';

@Entity()
export class Tickets {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  type: string;

  @Column()
  availableUnits: string;

  @Column()
  name: string;

  @Column()
  description: string;

  @Column()
  createdAt: Date;

  @Column()
  updatedAt: Date;
}
