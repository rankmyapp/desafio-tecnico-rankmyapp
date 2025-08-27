import {
  Body,
  Controller,
  Delete,
  Get,
  Logger,
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

@Controller({ path: 'tickets', version: '1' })
export class TicketsController {
  private readonly logger = new Logger(TicketsController.name);

  constructor(private ticketsService: TicketsService) {}

  @Post()
  @UseGuards(AuthGuard)
  create(@Body() body: CreateTicketDto) {
    this.logger.log(
      `Received request to create ticket with type: ${body.type}`,
    );
    return this.ticketsService.create(body);
  }

  @Get('catalog')
  @UseGuards(AuthGuard)
  getCatalog() {
    this.logger.log('Received request to get ticket catalog with available stock');
    return this.ticketsService.find();
  }

  @Get()
  @UseGuards(AuthGuard)
  findAll() {
    this.logger.log('Received request to find all tickets');
    return this.ticketsService.find();
  }

  @Get('/:id')
  @UseGuards(AuthGuard)
  findOne(@Param('id') id: string) {
    this.logger.log(`Received request to find ticket with ID: ${id}`);
    return this.ticketsService.findOne(parseInt(id));
  }

  @Patch('/:id')
  @UseGuards(AuthGuard)
  update(@Param('id') id: string, @Body() body: CreateTicketDto) {
    this.logger.log(`Received request to update ticket with ID: ${id}`);
    this.logger.debug(`Update data: ${JSON.stringify(body)}`);
    return this.ticketsService.update(parseInt(id), body);
  }

  @Delete('/:id')
  @UseGuards(AuthGuard)
  remove(@Param('id') id: string) {
    this.logger.log(`Received request to remove ticket with ID: ${id}`);
    return this.ticketsService.remove(parseInt(id));
  }

  @Post('buy')
  @UseGuards(AuthGuard)
  buy(@Body() body: BuyTicketDto, @Req() req) {
    const userId = req.user.sub as number;
    this.logger.log(
      `Received request to buy ticket ID: ${body.ticketId} for user ID: ${body.userId}`,
    );
    this.logger.debug(`Request made by authenticated user ID: ${userId}`);

    return this.ticketsService.buy({
      reqUserId: userId,
      ticketId: body.ticketId,
      userId: body.userId,
    });
  }
}
