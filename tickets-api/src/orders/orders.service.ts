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

  create(params: { ticketId: number; status: string; userId: number }) {
    const { ticketId, status, userId } = params;
    const order = this.repo.create({
      ticketId,
      status,
      userId,
    });

    return this.repo.save(order);
  }

  findOne(id: number) {
    return this.repo.findOne({ 
      where: { id },
      relations: ['user', 'ticket']
    });
  }

  find(filters: Partial<Order> = {}) {
    return this.repo.find({ 
      where: filters,
      relations: ['user', 'ticket']
    });
  }
  
  findByUserId(userId: number) {
    return this.repo.find({
      where: { userId },
      relations: ['user', 'ticket']
    });
  }
  
  countByUserId(userId: number) {
    return this.repo.count({
      where: { userId }
    });
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
