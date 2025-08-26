import { Equals, IsNumber } from 'class-validator';

export class BuyTicketDto {
  @IsNumber()
  ticketId: number;

  @Equals('CREDIT_CARD', {
    message: 'paymentType must be CREDIT_CARD',
  })
  paymentType: string;

  @IsNumber()
  userId: number;
}
