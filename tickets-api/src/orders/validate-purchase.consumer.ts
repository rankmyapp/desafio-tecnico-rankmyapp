import { Processor, WorkerHost } from '@nestjs/bullmq';
import { Job } from 'bullmq';
import { OrdersService } from './orders.service';
import {
  BadRequestException,
  NotFoundException,
  InternalServerErrorException,
  Logger,
} from '@nestjs/common';
import { Ticket } from '../tickets/tickets.entity';
import { Repository } from 'typeorm';
import { InjectRepository } from '@nestjs/typeorm';

const sleep = (ms: number) => new Promise((res) => setTimeout(res, ms));

@Processor('validate-purchase')
export class ValidatePurchaseConsumer extends WorkerHost {
  private readonly logger = new Logger(ValidatePurchaseConsumer.name);

  constructor(
    private readonly ordersService: OrdersService,
    @InjectRepository(Ticket) private repo: Repository<Ticket>,
  ) {
    super();
    this.logger.log('Validate Purchase Consumer initialized');
  }

  async process(job: Job<any, any, string>): Promise<any> {
    this.logger.log(
      `Processing job: ${job.name} with data: ${JSON.stringify(job.data)}`,
    );

    if (job.name === 'buy-ticket') {
      this.logger.debug('Validating job data structure');
      if (!job.data.orderId) {
        this.logger.error(
          `Invalid job data: Order ID not provided in ${JSON.stringify(job.data)}`,
        );
        throw new BadRequestException(`data not provided: ${job.data}`);
      }
      this.logger.debug('Job data structure is valid');

      const inComingOrderId = job.data.orderId as number;
      this.logger.debug(`Processing order ID: ${inComingOrderId}`);

      this.logger.debug(`Finding order with ID: ${inComingOrderId}`);
      const order = await this.ordersService.findOne(inComingOrderId);
      if (!order) {
        this.logger.warn(
          `Order with ID: ${inComingOrderId} not found - processing failed`,
        );
        throw new BadRequestException('Order not found');
      }
      this.logger.debug(
        `Found order with ID: ${inComingOrderId}, ticket ID: ${order.ticketId}`,
      );

      this.logger.debug(`Finding ticket with ID: ${order.ticketId}`);
      const ticket = await this.repo.findOneBy({
        id: order.ticketId,
      });

      if (!ticket) {
        this.logger.warn(
          `Ticket with ID: ${order.ticketId} not found - processing failed`,
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
          `No available units for ticket ID: ${ticket.id} - processing failed`,
        );
        throw new InternalServerErrorException('No tickets available!');
      }
      this.logger.debug('Ticket availability confirmed');

      this.logger.log('Simulating payment processing delay');
      await sleep(5000);

      this.logger.debug(
        `Decrementing available units for ticket ID: ${ticket.id} from ${ticket.availableUnits} to ${ticket.availableUnits - 1}`,
      );
      ticket.availableUnits -= 1;
      await this.repo.save(ticket);
      this.logger.debug(`Ticket ${ticket.id} inventory updated successfully`);

      this.logger.debug(`Updating order ${inComingOrderId} status to 'paid'`);
      await this.ordersService.update(inComingOrderId, {
        status: 'paid',
      });

      this.logger.log(
        `Order ${inComingOrderId} payment processed successfully, ticket ${ticket.id} now has ${ticket.availableUnits} units available`,
      );
    }

    return {};
  }
}
