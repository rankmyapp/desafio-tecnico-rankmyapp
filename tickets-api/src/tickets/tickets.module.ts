import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Tickets } from './tickets.entity';
import { TicketsController } from './tickets.controller';
import { TicketsService } from './tickets.service';

@Module({
  imports: [TypeOrmModule.forFeature([Tickets])],
  controllers: [TicketsController],
  providers: [TicketsService],
})
export class TicketsModule {}
