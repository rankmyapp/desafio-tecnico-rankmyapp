import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UsersModule } from './users/users.module';
import { TicketsModule } from './tickets/tickets.module';
import { User } from './users/users.entity';
import { Ticket } from './tickets/tickets.entity';
import { Order } from './orders/orders.entity';
import { OrdersModule } from './orders/orders.module';
import { JwtModule } from '@nestjs/jwt';
import { BullQueueModule } from './bull/bull.module';


@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      envFilePath: '.env',
    }),
    JwtModule.registerAsync({
      global: true,
      inject: [ConfigService],
      useFactory: (configService: ConfigService) => ({
        secret: configService.get<string>('JWT_SECRET'),
        signOptions: { expiresIn: configService.get<string>('JWT_EXPIRATION', '1h') },
      }),
    }),
    TypeOrmModule.forRootAsync({
      inject: [ConfigService],
      useFactory: (configService: ConfigService) => ({
        type: configService.get<any>('DB_TYPE', 'mysql'),
        host: configService.get<string>('DB_HOST', 'mysql_db'),
        port: configService.get<number>('DB_PORT', 3306),
        username: configService.get<string>('DB_USERNAME', 'daniel'),
        password: configService.get<string>('DB_PASSWORD', 'daniel'),
        database: configService.get<string>('DB_DATABASE', 'mydb'),
        entities: [User, Ticket, Order],
        synchronize: configService.get<boolean>('DB_SYNCHRONIZE', true),
      }),
    }),
    UsersModule,
    TicketsModule,
    OrdersModule,
    BullQueueModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
