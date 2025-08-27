import { Processor, WorkerHost } from '@nestjs/bullmq';
import { Job } from 'bullmq';
import { OrdersService } from './orders.service';
import {
  BadRequestException,
  NotFoundException,
  InternalServerErrorException,
} from '@nestjs/common';
import { Ticket } from '../tickets/tickets.entity';
import { Repository } from 'typeorm';
import { InjectRepository } from '@nestjs/typeorm';

const sleep = (ms: number) => new Promise((res) => setTimeout(res, ms));

@Processor('validate-purchase')
export class ValidatePurchaseConsumer extends WorkerHost {
  constructor(
    private readonly ordersService: OrdersService,
    @InjectRepository(Ticket) private repo: Repository<Ticket>,
  ) {
    super();
  }

  async process(job: Job<any, any, string>): Promise<any> {
    console.log('Processing job:', job.name, job.data);

    if (job.name === 'buy-ticket') {
      if (!job.data.orderId) {
        throw new BadRequestException(`data not provided: ${job.data}`);
      }

      const inComingOrderId = job.data.orderId as number;

      const order = await this.ordersService.findOne(inComingOrderId);
      if (!order) {
        throw new BadRequestException('Order not found');
      }

      const ticket = await this.repo.findOneBy({
        id: order.ticketId,
      });

      if (!ticket) {
        throw new NotFoundException('Ticket not found!');
      }

      if (!ticket.availableUnits) {
        throw new InternalServerErrorException('No tickets available!');
      }

      console.log('Simulando delay');
      await sleep(5000);

      ticket.availableUnits -= 1;
      await this.repo.save(ticket);

      await this.ordersService.update(inComingOrderId, {
        status: 'paid',
      });

      console.log('Agora ta pago');
    }

    return {};
  }
}
