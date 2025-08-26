import {
  Injectable,
  InternalServerErrorException,
  NotFoundException,
} from '@nestjs/common';
import { Repository } from 'typeorm';
import { Ticket } from './tickets.entity';
import { InjectRepository } from '@nestjs/typeorm';
import { OrdersService } from 'src/orders/orders.service';

@Injectable()
export class TicketsService {
  constructor(
    @InjectRepository(Ticket) private repo: Repository<Ticket>,
    private ordersService: OrdersService,
  ) {}

  create(params: {
    type: string;
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

  async buy(params: { ticketId: number; userId: number }) {
    const { ticketId } = params;

    const ticket = await this.repo.findOneBy({
      id: ticketId,
    });

    if (!ticket) {
      throw new NotFoundException('Ticket not found!');
    }

    if (!ticket.availableUnits) {
      throw new InternalServerErrorException('No tickets available!');
    }

    await this.ordersService.create({
      status: 'pendingPayment',
      ticketId: ticket.id,
    });

    // Call the producer Here

    return;
  }
}
