import { IsNumber, IsString } from 'class-validator';

export class CreateTicketDto {
  @IsString()
  type: string;

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
