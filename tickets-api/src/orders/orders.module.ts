import { Module } from '@nestjs/common';
import { OrdersController } from './orders.controller';
import { OrdersService } from './orders.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { ValidatePurchaseConsumer } from './validate-purchase.consumer';

import { Order } from './orders.entity';
import { Ticket } from '../tickets/tickets.entity';

@Module({
  imports: [TypeOrmModule.forFeature([Order, Ticket])],
  controllers: [OrdersController],
  providers: [OrdersService, ValidatePurchaseConsumer],
  exports: [OrdersService],
})
export class OrdersModule {}
