import { IsEnum, IsNumber, IsOptional, IsString } from 'class-validator';
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

  @IsOptional()
  @IsString()
  name?: string;

  @IsOptional()
  @IsString()
  description?: string;
}
