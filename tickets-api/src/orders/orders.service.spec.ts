import { Test, TestingModule } from '@nestjs/testing';
import { OrdersService } from './orders.service';
import { getRepositoryToken } from '@nestjs/typeorm';
import { Order } from './orders.entity';
import { Repository } from 'typeorm';
import { NotFoundException } from '@nestjs/common';

describe('OrdersService', () => {
  let service: OrdersService;
  let repository: jest.Mocked<Repository<Order>>;

  const mockRepository = () => ({
    create: jest.fn(),
    save: jest.fn(),
    findOne: jest.fn(),
    find: jest.fn(),
    count: jest.fn(),
    remove: jest.fn(),
  });

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        OrdersService,
        {
          provide: getRepositoryToken(Order),
          useFactory: mockRepository,
        },
      ],
    }).compile();

    service = module.get<OrdersService>(OrdersService);
    repository = module.get(getRepositoryToken(Order));
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  describe('create', () => {
    it('should create and save a new order', async () => {
      const orderData = {
        ticketId: 1,
        status: 'pendingPayment',
        userId: 2,
      };
      const createdOrder = { ...orderData, id: 1 } as Order;

      repository.create.mockReturnValue(createdOrder);
      repository.save.mockResolvedValue(createdOrder);

      const result = await service.create(orderData);

      expect(repository.create).toHaveBeenCalledWith(orderData);

      expect(repository.save).toHaveBeenCalledWith(createdOrder);
      expect(result).toEqual(createdOrder);
    });
  });

  describe('findOne', () => {
    it('should return an order if found', async () => {
      const mockOrder = {
        id: 1,
        ticketId: 1,
        userId: 2,
        status: 'paid',
      } as Order;
      repository.findOne.mockResolvedValue(mockOrder);

      const result = await service.findOne(1);

      expect(repository.findOne).toHaveBeenCalledWith({
        where: { id: 1 },
        relations: ['user', 'ticket'],
      });
      expect(result).toEqual(mockOrder);
    });

    it('should throw NotFoundException if order is not found', async () => {
      repository.findOne.mockResolvedValue(null);

      await expect(service.findOne(999)).rejects.toThrow(
        new NotFoundException('No order found with ID: 999'),
      );

      expect(repository.findOne).toHaveBeenCalledWith({
        where: { id: 999 },
        relations: ['user', 'ticket'],
      });
    });
  });

  describe('find', () => {
    it('should return all orders when no filters are provided', async () => {
      const mockOrders = [
        { id: 1, ticketId: 1, userId: 2, status: 'paid' },
        { id: 2, ticketId: 3, userId: 2, status: 'pendingPayment' },
      ] as Order[];
      repository.find.mockResolvedValue(mockOrders);

      const result = await service.find();

      expect(repository.find).toHaveBeenCalledWith({
        where: {},
        relations: ['user', 'ticket'],
      });
      expect(result).toEqual(mockOrders);
    });

    it('should return filtered orders when filters are provided', async () => {
      const filters = { status: 'paid' };
      const mockOrders = [
        { id: 1, ticketId: 1, userId: 2, status: 'paid' },
      ] as Order[];
      repository.find.mockResolvedValue(mockOrders);

      const result = await service.find(filters);

      expect(repository.find).toHaveBeenCalledWith({
        where: filters,
        relations: ['user', 'ticket'],
      });
      expect(result).toEqual(mockOrders);
    });
  });

  describe('findByUserId', () => {
    it('should return orders for a specific user', async () => {
      const userId = 2;
      const mockOrders = [
        { id: 1, ticketId: 1, userId, status: 'paid' },
        { id: 2, ticketId: 3, userId, status: 'pendingPayment' },
      ] as Order[];
      repository.find.mockResolvedValue(mockOrders);

      const result = await service.findByUserId(userId);

      expect(repository.find).toHaveBeenCalledWith({
        where: { userId },
        relations: ['user', 'ticket'],
      });
      expect(result).toEqual(mockOrders);
    });
  });

  describe('countByUserId', () => {
    it('should return the count of orders for a specific user', async () => {
      const userId = 2;
      repository.count.mockResolvedValue(5);

      const result = await service.countByUserId(userId);

      expect(repository.count).toHaveBeenCalledWith({
        where: { userId },
      });
      expect(result).toEqual(5);
    });
  });

  describe('update', () => {
    it('should update an order if found', async () => {
      const orderId = 1;
      const updateData = { status: 'paid' };
      const existingOrder = {
        id: orderId,
        ticketId: 1,
        userId: 2,
        status: 'pendingPayment',
      } as Order;
      const updatedOrder = { ...existingOrder, ...updateData };

      repository.findOne.mockResolvedValue(existingOrder);
      repository.save.mockResolvedValue(updatedOrder);

      const result = await service.update(orderId, updateData);

      expect(repository.findOne).toHaveBeenCalledWith({
        where: { id: orderId },
        relations: ['user', 'ticket'],
      });
      expect(repository.save).toHaveBeenCalledWith(updatedOrder);
      expect(result).toEqual(updatedOrder);
    });

    it('should throw NotFoundException if order is not found', async () => {
      repository.findOne.mockResolvedValue(null);

      await expect(service.update(999, { status: 'paid' })).rejects.toThrow(
        NotFoundException,
      );
      expect(repository.findOne).toHaveBeenCalledWith({
        where: { id: 999 },
        relations: ['user', 'ticket'],
      });
      expect(repository.save).not.toHaveBeenCalled();
    });
  });

  describe('remove', () => {
    it('should remove an order if found', async () => {
      const orderId = 1;
      const existingOrder = {
        id: orderId,
        ticketId: 1,
        userId: 2,
        status: 'pendingPayment',
      } as Order;

      repository.findOne.mockResolvedValue(existingOrder);
      repository.remove.mockResolvedValue(existingOrder);

      const result = await service.remove(orderId);

      expect(repository.findOne).toHaveBeenCalledWith({
        where: { id: orderId },
        relations: ['user', 'ticket'],
      });
      expect(repository.remove).toHaveBeenCalledWith(existingOrder);
      expect(result).toEqual(existingOrder);
    });

    it('should throw NotFoundException if order is not found', async () => {
      repository.findOne.mockResolvedValue(null);

      await expect(service.remove(999)).rejects.toThrow(NotFoundException);
      expect(repository.findOne).toHaveBeenCalledWith({
        where: { id: 999 },
        relations: ['user', 'ticket'],
      });
      expect(repository.remove).not.toHaveBeenCalled();
    });
  });
});
