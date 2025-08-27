import {
  BadRequestException,
  Injectable,
  NotFoundException,
  Logger,
} from '@nestjs/common';
import { UsersService } from '../users.service';
import { randomBytes, scrypt as _scrypt } from 'crypto';
import { promisify } from 'util';
import { JwtService } from '@nestjs/jwt';
import { ConfigService } from '@nestjs/config';

const scrypt = promisify(_scrypt);

@Injectable()
export class AuthService {
  private readonly logger = new Logger(AuthService.name);

  constructor(
    private usersService: UsersService,
    private jwtService: JwtService,
    private configService: ConfigService,
  ) {
    this.logger.log('Auth Service initialized');
  }

  async signUp(params: { email: string; password: string }) {
    const { email, password } = params;
    const users = await this.usersService.find(email);
    if (users.length) {
      throw new BadRequestException('Email already exists');
    }

    const salt = randomBytes(8).toString('hex');
    const hash = (await scrypt(password, salt, 32)) as Buffer;
    const result = salt + '.' + hash.toString('hex');

    this.usersService.create({
      email,
      password: result,
    });
  }

  async signIn(params: { email: string; password: string }) {
    const { email, password } = params;
    const [user] = await this.usersService.find(email);
    if (!user) {
      throw new NotFoundException('Email not found!');
    }

    const [salt, storedHash] = user.password.split('.');
    const hash = (await scrypt(password, salt, 32)) as Buffer;

    if (storedHash !== hash.toString('hex')) {
      throw new BadRequestException('Failed to login');
    }

    const payload = { sub: user.id, email: user.email };
    const jwtExpiration = this.configService.get<string>(
      'JWT_EXPIRATION',
      '1h',
    );

    this.logger.debug(
      `Generating JWT token for user ID: ${user.id} with expiration: ${jwtExpiration}`,
    );
    return {
      access_token: await this.jwtService.signAsync(payload, {
        expiresIn: jwtExpiration,
      }),
    };
  }
}
