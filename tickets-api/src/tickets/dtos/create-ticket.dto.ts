import { IsEnum, IsNumber, IsString } from 'class-validator';
import { TicketsType } from '../tickets.entity';

export class CreateTicketDto {
  @IsEnum(TicketsType, {
    message: 'Type must be one of: generalArea, grandStand, vip, goldenCircle',
  })
  type: TicketsType;

  @IsNumber()
  availableUnits: number;

  @IsNumber()
  price: number;

  @IsString({
    always: false,
  })
  name?: string;

  @IsString({
    always: false,
  })
  description?: string;
}
