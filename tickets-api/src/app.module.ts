import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UsersModule } from './users/users.module';
import { TicketsController } from './tickets/tickets.controller';
import { TicketsService } from './tickets/tickets.service';
import { TicketsModule } from './tickets/tickets.module';
import { User } from './users/users.entity';
import { Tickets } from './tickets/tickets.entity';

@Module({
  imports: [
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: 'mysql_db',
      database: 'database',
      entities: [User, Tickets],
      synchronize: true,
      username: 'daniel',
      password: '123',
    }),
    UsersModule,
    TicketsModule,
  ],
  controllers: [AppController, TicketsController],
  providers: [AppService, TicketsService],
})
export class AppModule {}
