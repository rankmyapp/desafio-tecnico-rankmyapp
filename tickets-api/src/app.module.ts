import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UsersModule } from './users/users.module';
import { TicketsModule } from './tickets/tickets.module';
import { User } from './users/users.entity';
import { Ticket } from './tickets/tickets.entity';
import { Order } from './orders/orders.entity';
import { OrdersModule } from './orders/orders.module';
import { JwtModule } from '@nestjs/jwt';
import { jwtConstants } from './users/auth/constants';
import { BullQueueModule } from './bull/bull.module';


@Module({
  imports: [
    JwtModule.register({
      global: true,
      secret: jwtConstants.secret,
      signOptions: { expiresIn: '60s' },
    }),
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: 'mysql_db',
      database: 'mydb',
      entities: [User, Ticket, Order],
      synchronize: true,
      username: 'daniel',
      password: 'daniel',
    }),
    UsersModule,
    TicketsModule,
    OrdersModule,
    BullQueueModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
