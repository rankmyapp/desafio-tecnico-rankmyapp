import {
  Body,
  Controller,
  Delete,
  Get,
  Param,
  Patch,
  Post,
} from '@nestjs/common';
import { TicketsService } from './tickets.service';
import { CreateTicketDto } from './dtos/create-ticket.dto';
import { BuyTicketDto } from './dtos/buy-ticket.dto';

@Controller('tickets')
export class TicketsController {
  constructor(private ticketsService: TicketsService) {}

  @Post()
  create(@Body() body: CreateTicketDto) {
    return this.ticketsService.create(body);
  }

  @Get()
  findAll() {
    return this.ticketsService.find();
  }

  @Get('/:id')
  findOne(@Param('id') id: string) {
    return this.ticketsService.findOne(parseInt(id));
  }

  @Patch('/:id')
  update(@Param('id') id: string, @Body() body: CreateTicketDto) {
    return this.ticketsService.update(parseInt(id), body);
  }

  @Delete('/:id')
  remove(@Param('id') id: string) {
    return this.ticketsService.remove(parseInt(id));
  }

  @Post('/buy')
  buy(@Body() body: BuyTicketDto) {
    return this.ticketsService.buy(body);
  }
}
