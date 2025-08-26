import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Ticket } from './tickets.entity';
import { TicketsController } from './tickets.controller';
import { TicketsService } from './tickets.service';
import { OrdersModule } from 'src/orders/orders.module';
import { BullQueueModule } from 'src/bull/bull.module';

@Module({
  imports: [TypeOrmModule.forFeature([Ticket]), OrdersModule, BullQueueModule],
  controllers: [TicketsController],
  providers: [TicketsService],
  exports: [TicketsService],
})
export class TicketsModule {}
