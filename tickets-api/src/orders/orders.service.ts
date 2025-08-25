import { Injectable, NotFoundException } from '@nestjs/common';
import { Repository } from 'typeorm';
import { Order } from './orders.entity';
import { InjectRepository } from '@nestjs/typeorm';

@Injectable()
export class OrdersService {
  private repo: Repository<Order>;

  constructor(@InjectRepository(Order) repo: Repository<Order>) {
    this.repo = repo;
  }

  create(params: { ticketId: number; status: string }) {
    const { ticketId, status } = params;
    const order = this.repo.create({
      ticketId,
      status,
    });

    return this.repo.save(order);
  }

  findOne(id: number) {
    return this.repo.findOneBy({ id });
  }

  find(filters: Partial<Order> = {}) {
    return this.repo.find({ where: filters });
  }

  async update(id: number, params: Partial<Order>) {
    const order = await this.findOne(id);
    if (!order) {
      throw new NotFoundException('Order not found');
    }

    Object.assign(order, params);
    return this.repo.save(order);
  }

  async remove(id: number) {
    const order = await this.findOne(id);
    if (!order) {
      throw new NotFoundException('Order not found');
    }

    return this.repo.remove(order);
  }
}
