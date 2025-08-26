import { Processor, WorkerHost } from '@nestjs/bullmq';
import { Job } from 'bullmq';
import { OrdersService } from './orders.service';
import { BadRequestException } from '@nestjs/common';

const sleep = (ms: number) => new Promise((res) => setTimeout(res, ms));

@Processor('process-orders')
export class ProcessOrderConsumer extends WorkerHost {
  constructor(private readonly ordersService: OrdersService) {
    super();
  }

  async process(job: Job<any, any, string>): Promise<any> {
    console.log('Processing job:', job.name, job.data);

    if (job.name === 'buy-ticket') {
      if (!job.data.ticketId || !job.data.userId) {
        throw new BadRequestException(`data not provided: ${job.data}`);
      }
      console.log('Simulando delay');
      await sleep(2000);

      await this.ordersService.create({
        status: 'paid',
        ticketId: job.data.ticketId as number,
        userId: job.data.userId as number,
      });

      console.log('Agora ta pago');
    }

    return {};
  }
}
