import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UsersModule } from './users/users.module';
import { TicketsController } from './tickets/tickets.controller';
import { TicketsService } from './tickets/tickets.service';
import { TicketsModule } from './tickets/tickets.module';
import { User } from './users/users.entity';
import { Ticket } from './tickets/tickets.entity';
import { Order } from './orders/orders.entity';
import { OrdersModule } from './orders/orders.module';
import { OrdersController } from './orders/orders.controller';
import { OrdersService } from './orders/orders.service';

@Module({
  imports: [
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: 'mysql_db',
      database: 'database',
      entities: [User, Ticket, Order],
      synchronize: true,
      username: 'daniel',
      password: '123',
    }),
    UsersModule,
    TicketsModule,
    OrdersModule,
  ],
  controllers: [AppController, TicketsController, OrdersController],
  providers: [AppService, TicketsService, OrdersService],
})
export class AppModule {}
