import {
  BadRequestException,
  Injectable,
  InternalServerErrorException,
  NotFoundException,
} from '@nestjs/common';
import { Repository } from 'typeorm';
import { Ticket, TicketsType } from './tickets.entity';
import { InjectRepository } from '@nestjs/typeorm';
import { OrdersService } from '../orders/orders.service';
import { InjectQueue } from '@nestjs/bullmq';
import { Queue } from 'bullmq';

@Injectable()
export class TicketsService {
  constructor(
    @InjectRepository(Ticket) private repo: Repository<Ticket>,
    private ordersService: OrdersService,
    @InjectQueue('validate-purchase') private processOrdersQueue: Queue,
  ) {}

  create(params: {
    type: TicketsType;
    availableUnits: number;
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

  async buy(params: { ticketId: number; userId: number; reqUserId: number }) {
    const { ticketId, userId, reqUserId } = params;

    console.log('INcoming reqUserId', reqUserId);

    if (ticketId !== reqUserId) {
      throw new BadRequestException('The user id is not the same');
    }

    const ticket = await this.repo.findOneBy({
      id: ticketId,
    });

    if (!ticket) {
      throw new NotFoundException('Ticket not found!');
    }

    if (!ticket.availableUnits) {
      throw new InternalServerErrorException('No tickets available!');
    }

    const order = await this.ordersService.create({
      status: 'pendingPayment',
      ticketId: ticketId,
      userId: userId,
    });

    await this.processOrdersQueue.add('buy-ticket', {
      orderId: order.id,
    });

    return {
      status: order.status,
    };
  }
}
