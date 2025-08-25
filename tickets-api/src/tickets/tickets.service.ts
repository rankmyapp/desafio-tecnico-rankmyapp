import { Injectable, NotFoundException } from '@nestjs/common';
import { Repository } from 'typeorm';
import { Ticket } from './tickets.entity';
import { InjectRepository } from '@nestjs/typeorm';

@Injectable()
export class TicketsService {
  private repo: Repository<Ticket>;

  constructor(@InjectRepository(Ticket) repo: Repository<Ticket>) {
    this.repo = repo;
  }

  create(params: {
    type: string;
    availableUnits: string;
    price: number;
    name?: string;
    description?: string;
  }) {
    const { type, availableUnits, price, name, description } = params;
    const ticket = this.repo.create({
      type,
      availableUnits,
      price,
      name,
      description,
    });

    return this.repo.save(ticket);
  }

  findOne(id: number) {
    return this.repo.findOneBy({ id });
  }

  find(filters: Partial<Ticket> = {}) {
    return this.repo.find({ where: filters });
  }

  async update(id: number, params: Partial<Ticket>) {
    const ticket = await this.findOne(id);
    if (!ticket) {
      throw new NotFoundException('Ticket not found');
    }

    Object.assign(ticket, params);
    return this.repo.save(ticket);
  }

  async remove(id: number) {
    const ticket = await this.findOne(id);
    if (!ticket) {
      throw new NotFoundException('Ticket not found');
    }

    return this.repo.remove(ticket);
  }
}
