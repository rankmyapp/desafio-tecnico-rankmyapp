import { Injectable, Logger, NotFoundException } from '@nestjs/common';
import { Repository } from 'typeorm';
import { Order } from './orders.entity';
import { InjectRepository } from '@nestjs/typeorm';

@Injectable()
export class OrdersService {
  private readonly logger = new Logger(OrdersService.name);
  private repo: Repository<Order>;

  constructor(@InjectRepository(Order) repo: Repository<Order>) {
    this.repo = repo;
    this.logger.log('Orders Service initialized');
  }

  async create(params: { ticketId: number; status: string; userId: number }) {
    this.logger.log(
      `Creating order for ticket ID: ${params.ticketId}, user ID: ${params.userId}, status: ${params.status}`,
    );
    const { ticketId, status, userId } = params;

    this.logger.debug('Creating order entity');
    const order = this.repo.create({
      ticketId,
      status,
      userId,
    });

    this.logger.debug('Saving order to database');
    const savedOrder = await this.repo.save(order);
    this.logger.log(`Order created successfully with ID: ${savedOrder.id}`);
    return savedOrder;
  }

  async findOne(id: number) {
    this.logger.debug(`Finding order with ID: ${id}`);
    const order = await this.repo.findOne({
      where: { id },
      relations: ['user', 'ticket'],
    });

    if (!order) {
      this.logger.debug(`No order found with ID: ${id}`);
      throw new NotFoundException(`No order found with ID: ${id}`);
    }

    this.logger.debug(`Found order with ID: ${id}`);

    return order;
  }

  async find(filters: Partial<Order> = {}) {
    this.logger.debug(
      `Finding orders with filters: ${JSON.stringify(filters)}`,
    );
    const orders = await this.repo.find({
      where: filters,
      relations: ['user', 'ticket'],
    });

    this.logger.debug(`Found ${orders.length} orders matching filters`);
    return orders;
  }

  async findByUserId(userId: number) {
    this.logger.debug(`Finding orders for user ID: ${userId}`);
    const orders = await this.repo.find({
      where: { userId },
      relations: ['user', 'ticket'],
    });

    this.logger.debug(`Found ${orders.length} orders for user ID: ${userId}`);
    return orders;
  }

  async countByUserId(userId: number) {
    this.logger.debug(`Counting orders for user ID: ${userId}`);
    const count = await this.repo.count({
      where: { userId },
    });

    this.logger.debug(`User ID: ${userId} has ${count} orders`);
    return count;
  }

  async update(id: number, params: Partial<Order>) {
    this.logger.log(`Updating order with ID: ${id}`);
    this.logger.debug(`Update parameters: ${JSON.stringify(params)}`);

    const order = await this.findOne(id);
    if (!order) {
      this.logger.warn(`Order with ID: ${id} not found - update failed`);
      throw new NotFoundException('Order not found');
    }

    this.logger.debug(`Applying updates to order ${id}`);
    Object.assign(order, params);

    this.logger.debug('Saving updated order to database');
    const updatedOrder = await this.repo.save(order);
    this.logger.log(`Order ${id} updated successfully`);
    return updatedOrder;
  }

  async remove(id: number) {
    this.logger.log(`Removing order with ID: ${id}`);

    const order = await this.findOne(id);
    if (!order) {
      this.logger.warn(`Order with ID: ${id} not found - removal failed`);
      throw new NotFoundException('Order not found');
    }

    this.logger.debug(`Removing order ${id} from database`);
    const removedOrder = await this.repo.remove(order);
    this.logger.log(`Order ${id} removed successfully`);
    return removedOrder;
  }
}
