import { Module } from '@nestjs/common';
import { OrdersController } from './orders.controller';
import { OrdersService } from './orders.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { ProcessOrderConsumer } from './process-order.consumer';

import { Order } from './orders.entity';

@Module({
  imports: [TypeOrmModule.forFeature([Order])],
  controllers: [OrdersController],
  providers: [OrdersService, ProcessOrderConsumer],
  exports: [OrdersService]
})
export class OrdersModule {}
