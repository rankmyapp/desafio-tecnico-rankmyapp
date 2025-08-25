import { IsString } from 'class-validator';

export class CreateTicketDto {
  @IsString()
  type: string;

  @IsString()
  availableUnits: string;

  @IsString()
  price: number;

  @IsString()
  name?: string;

  @IsString()
  description?: string;
}
