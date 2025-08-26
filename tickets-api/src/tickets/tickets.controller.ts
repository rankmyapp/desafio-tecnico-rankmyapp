import {
  Body,
  Controller,
  Delete,
  Get,
  Param,
  Patch,
  Post,
  Req,
  UseGuards,
} from '@nestjs/common';
import { TicketsService } from './tickets.service';
import { CreateTicketDto } from './dtos/create-ticket.dto';
import { BuyTicketDto } from './dtos/buy-ticket.dto';
import { AuthGuard } from 'src/users/auth/auth.guard';

@Controller('tickets')
export class TicketsController {
  constructor(private ticketsService: TicketsService) {}

  @Post()
  @UseGuards(AuthGuard)
  create(@Body() body: CreateTicketDto) {
    return this.ticketsService.create(body);
  }

  @Get()
  @UseGuards(AuthGuard)
  findAll() {
    return this.ticketsService.find();
  }

  @Get('/:id')
  @UseGuards(AuthGuard)
  findOne(@Param('id') id: string) {
    return this.ticketsService.findOne(parseInt(id));
  }

  @Patch('/:id')
  @UseGuards(AuthGuard)
  update(@Param('id') id: string, @Body() body: CreateTicketDto) {
    return this.ticketsService.update(parseInt(id), body);
  }

  @Delete('/:id')
  @UseGuards(AuthGuard)
  remove(@Param('id') id: string) {
    return this.ticketsService.remove(parseInt(id));
  }

  @Post('/buy')
  @UseGuards(AuthGuard)
  buy(@Body() body: BuyTicketDto, @Req() req) {
    const userId = req.user.sub as number;

    return this.ticketsService.buy({
      reqUserId: userId,
      ticketId: body.ticketId,
      userId: body.userId,
    });
  }
}
