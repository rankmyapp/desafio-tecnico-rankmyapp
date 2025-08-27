import {
  BadRequestException,
  Injectable,
  InternalServerErrorException,
  NotFoundException,
  Logger,
} from '@nestjs/common';
import { Repository } from 'typeorm';
import { Ticket, TicketsType } from './tickets.entity';
import { InjectRepository } from '@nestjs/typeorm';
import { OrdersService } from '../orders/orders.service';
import { InjectQueue } from '@nestjs/bullmq';
import { Queue } from 'bullmq';

@Injectable()
export class TicketsService {
  private readonly logger = new Logger(TicketsService.name);

  constructor(
    @InjectRepository(Ticket) private repo: Repository<Ticket>,
    private ordersService: OrdersService,
    @InjectQueue('validate-purchase') private processOrdersQueue: Queue,
  ) {}

  async create(params: {
    type: TicketsType;
    availableUnits: number;
    price: number;
    name?: string;
    description?: string;
  }) {
    this.logger.log(`Creating new ticket with type: ${params.type}`);
    const { type, availableUnits, price, name, description } = params;

    this.logger.debug(`Checking if ticket with type ${type} already exists`);
    const existingTicket = await this.repo.findOne({ where: { type } });
    if (existingTicket) {
      this.logger.warn(
        `Ticket with type ${type} already exists - creation rejected`,
      );
      throw new BadRequestException(`Ticket with type ${type} already exists`);
    }
    this.logger.debug('No duplicate ticket found, proceeding with creation');

    this.logger.debug(
      `Creating ticket entity with ${availableUnits} units at price ${price}`,
    );
    const ticket = this.repo.create({
      type,
      availableUnits,
      price,
      name,
      description,
    });

    this.logger.debug('Saving ticket to database');
    const savedTicket = await this.repo.save(ticket);
    this.logger.log(`Ticket created successfully with ID: ${savedTicket.id}`);
    return savedTicket;
  }

  async findOne(id: number) {
    this.logger.log(`Finding ticket with ID: ${id}`);
    const ticket = await this.repo.findOneBy({ id });
    if (ticket) {
      this.logger.debug(`Found ticket with ID: ${id}`);
    } else {
      this.logger.debug(`No ticket found with ID: ${id}`);
    }
    return ticket;
  }

  async find(filters: Partial<Ticket> = {}) {
    this.logger.log(`Finding tickets with filters: ${JSON.stringify(filters)}`);
    const tickets = await this.repo.find({ where: filters });
    this.logger.debug(`Found ${tickets.length} tickets matching filters`);
    return tickets;
  }

  async update(id: number, params: Partial<Ticket>) {
    this.logger.log(`Updating ticket with ID: ${id}`);
    this.logger.debug(`Update parameters: ${JSON.stringify(params)}`);

    const ticket = await this.findOne(id);
    if (!ticket) {
      this.logger.warn(`Ticket with ID: ${id} not found - update failed`);
      throw new NotFoundException('Ticket not found');
    }

    this.logger.debug(`Applying updates to ticket ${id}`);
    Object.assign(ticket, params);

    this.logger.debug('Saving updated ticket to database');
    const updatedTicket = await this.repo.save(ticket);
    this.logger.log(`Ticket ${id} updated successfully`);
    return updatedTicket;
  }

  async remove(id: number) {
    this.logger.log(`Removing ticket with ID: ${id}`);

    const ticket = await this.findOne(id);
    if (!ticket) {
      this.logger.warn(`Ticket with ID: ${id} not found - removal failed`);
      throw new NotFoundException('Ticket not found');
    }

    this.logger.debug(`Removing ticket ${id} from database`);
    const removedTicket = await this.repo.remove(ticket);
    this.logger.log(`Ticket ${id} removed successfully`);
    return removedTicket;
  }

  async buy(params: { ticketId: number; userId: number; reqUserId: number }) {
    const { ticketId, userId, reqUserId } = params;
    this.logger.log(
      `Processing purchase request for ticket ID: ${ticketId} by user ID: ${userId}`,
    );
    this.logger.debug(`Request user ID: ${reqUserId}`);

    this.logger.debug(
      `Validating user authorization (userId: ${userId}, reqUserId: ${reqUserId})`,
    );
    if (userId !== reqUserId) {
      this.logger.warn(
        `Authorization failed - userId (${userId}) doesn't match reqUserId (${reqUserId})`,
      );
      throw new BadRequestException('The user id is not the same');
    }
    this.logger.debug('User authorization validated successfully');

    this.logger.debug(`Finding ticket with ID: ${ticketId}`);
    const ticket = await this.findOne(ticketId);
    if (!ticket) {
      this.logger.warn(
        `Ticket with ID: ${ticketId} not found - purchase failed`,
      );
      throw new NotFoundException('Ticket not found!');
    }
    this.logger.debug(
      `Found ticket: ${ticket.type} with ${ticket.availableUnits} available units`,
    );

    this.logger.debug(
      `Checking ticket availability (available: ${ticket.availableUnits})`,
    );
    if (!ticket.availableUnits) {
      this.logger.warn(
        `No available units for ticket ID: ${ticketId} - purchase failed`,
      );
      throw new InternalServerErrorException('No tickets available!');
    }
    this.logger.debug('Ticket availability confirmed');

    this.logger.debug(
      `Creating order for ticket ID: ${ticketId} and user ID: ${userId}`,
    );
    const order = await this.ordersService.create({
      status: 'pendingPayment',
      ticketId: ticketId,
      userId: userId,
    });
    this.logger.debug(`Order created with ID: ${order.id}`);

    this.logger.debug(`Adding order ID: ${order.id} to processing queue`);
    await this.processOrdersQueue.add('buy-ticket', {
      orderId: order.id,
    });
    this.logger.log(
      `Purchase request for ticket ID: ${ticketId} processed successfully, order ID: ${order.id}`,
    );

    return {
      status: order.status,
    };
  }
}
