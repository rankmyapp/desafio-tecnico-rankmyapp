import { ConfigService } from '@nestjs/config';

export const jwtConstants = {
  secret: process.env.JWT_SECRET || 'anySecret',
};

export const getJwtConfig = (configService: ConfigService) => ({
  secret: configService.get<string>('JWT_SECRET'),
  expiresIn: configService.get<string>('JWT_EXPIRATION', '1h'),
});
