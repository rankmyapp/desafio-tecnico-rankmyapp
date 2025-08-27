import { Test, TestingModule } from '@nestjs/testing';
import { TicketsService } from './tickets.service';
import { getRepositoryToken } from '@nestjs/typeorm';
import { Ticket, TicketsType } from './tickets.entity';
import { Repository } from 'typeorm';
import {
  BadRequestException,
  InternalServerErrorException,
  NotFoundException,
} from '@nestjs/common';
import { OrdersService } from '../orders/orders.service';
import { getQueueToken } from '@nestjs/bullmq';
import { Queue } from 'bullmq';

describe('TicketsService', () => {
  let service: TicketsService;
  let repository: jest.Mocked<Repository<Ticket>>;
  let ordersService: jest.Mocked<OrdersService>;
  let processOrdersQueue: jest.Mocked<Queue>;

  const mockRepository = () => ({
    create: jest.fn(),
    save: jest.fn(),
    findOneBy: jest.fn(),
    findOne: jest.fn(),
    find: jest.fn(),
    remove: jest.fn(),
  });

  const mockOrdersService = () => ({
    create: jest.fn(),
  });

  const mockQueue = () => ({
    add: jest.fn(),
  });

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        TicketsService,
        {
          provide: getRepositoryToken(Ticket),
          useFactory: mockRepository,
        },
        {
          provide: OrdersService,
          useFactory: mockOrdersService,
        },
        {
          provide: getQueueToken('validate-purchase'),
          useFactory: mockQueue,
        },
      ],
    }).compile();

    service = module.get<TicketsService>(TicketsService);
    repository = module.get(getRepositoryToken(Ticket));
    ordersService = module.get(OrdersService);
    processOrdersQueue = module.get(getQueueToken('validate-purchase'));
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  describe('create', () => {
    it('should create and save a new ticket', async () => {
      const ticketData = {
        type: TicketsType.vip,
        availableUnits: 100,
        price: 50,
        name: 'Concert Ticket',
        description: 'A ticket for a concert',
      };
      const createdTicket = { ...ticketData, id: 1 } as Ticket;

      repository.findOne.mockResolvedValue(null);
      repository.create.mockReturnValue(createdTicket);
      repository.save.mockResolvedValue(createdTicket);

      const result = await service.create(ticketData);

      expect(repository.findOne).toHaveBeenCalledWith({
        where: { type: ticketData.type },
      });
      expect(repository.create).toHaveBeenCalledWith(ticketData);
      expect(repository.save).toHaveBeenCalledWith(createdTicket);
      expect(result).toEqual(createdTicket);
    });

    it('should throw BadRequestException when creating a ticket with a type that already exists', async () => {
      const ticketData = {
        type: TicketsType.vip,
        availableUnits: 100,
        price: 50,
      };
      const existingTicket = { ...ticketData, id: 1 } as Ticket;

      repository.findOne.mockResolvedValue(existingTicket);

      await expect(service.create(ticketData)).rejects.toThrow(
        BadRequestException,
      );
      expect(repository.findOne).toHaveBeenCalledWith({
        where: { type: ticketData.type },
      });
      expect(repository.create).not.toHaveBeenCalled();
      expect(repository.save).not.toHaveBeenCalled();
    });
  });

  describe('findOne', () => {
    it('should return a ticket if found', async () => {
      const mockTicket = {
        id: 1,
        type: TicketsType.vip,
        availableUnits: 100,
        price: 50,
      } as Ticket;
      repository.findOneBy.mockResolvedValue(mockTicket);

      const result = await service.findOne(1);

      expect(repository.findOneBy).toHaveBeenCalledWith({ id: 1 });
      expect(result).toEqual(mockTicket);
    });

    it('should return null if ticket is not found', async () => {
      repository.findOneBy.mockResolvedValue(null);

      const result = await service.findOne(999);

      expect(repository.findOneBy).toHaveBeenCalledWith({ id: 999 });
      expect(result).toBeNull();
    });
  });

  describe('find', () => {
    it('should return all tickets when no filters are provided', async () => {
      const mockTickets = [
        { id: 1, type: TicketsType.vip, availableUnits: 100, price: 50 },
        {
          id: 2,
          type: TicketsType.generalArea,
          availableUnits: 200,
          price: 15,
        },
      ] as Ticket[];
      repository.find.mockResolvedValue(mockTickets);

      const result = await service.find();

      expect(repository.find).toHaveBeenCalledWith({ where: {} });
      expect(result).toEqual(mockTickets);
    });

    it('should return filtered tickets when filters are provided', async () => {
      const filters = { type: TicketsType.vip };
      const mockTickets = [
        { id: 1, type: TicketsType.vip, availableUnits: 100, price: 50 },
      ] as Ticket[];
      repository.find.mockResolvedValue(mockTickets);

      const result = await service.find(filters);

      expect(repository.find).toHaveBeenCalledWith({ where: filters });
      expect(result).toEqual(mockTickets);
    });
  });

  describe('update', () => {
    it('should update a ticket if found', async () => {
      const ticketId = 1;
      const updateData = { price: 60 };
      const existingTicket = {
        id: ticketId,
        type: TicketsType.vip,
        availableUnits: 100,
        price: 50,
      } as Ticket;
      const updatedTicket = { ...existingTicket, ...updateData } as Ticket;

      repository.findOneBy.mockResolvedValue(existingTicket);
      repository.save.mockResolvedValue(updatedTicket);

      const result = await service.update(ticketId, updateData);

      expect(repository.findOneBy).toHaveBeenCalledWith({ id: ticketId });
      expect(repository.save).toHaveBeenCalledWith(updatedTicket);
      expect(result).toEqual(updatedTicket);
    });

    it('should throw NotFoundException if ticket is not found', async () => {
      repository.findOneBy.mockResolvedValue(null);

      await expect(service.update(999, { price: 60 })).rejects.toThrow(
        NotFoundException,
      );
      expect(repository.findOneBy).toHaveBeenCalledWith({ id: 999 });
      expect(repository.save).not.toHaveBeenCalled();
    });
  });

  describe('remove', () => {
    it('should remove a ticket if found', async () => {
      const ticketId = 1;
      const existingTicket = {
        id: ticketId,
        type: TicketsType.vip,
        availableUnits: 100,
        price: 50,
      } as Ticket;

      repository.findOneBy.mockResolvedValue(existingTicket);
      repository.remove.mockResolvedValue(existingTicket);

      const result = await service.remove(ticketId);

      expect(repository.findOneBy).toHaveBeenCalledWith({ id: ticketId });
      expect(repository.remove).toHaveBeenCalledWith(existingTicket);
      expect(result).toEqual(existingTicket);
    });

    it('should throw NotFoundException if ticket is not found', async () => {
      repository.findOneBy.mockResolvedValue(null);

      await expect(service.remove(999)).rejects.toThrow(NotFoundException);
      expect(repository.findOneBy).toHaveBeenCalledWith({ id: 999 });
      expect(repository.remove).not.toHaveBeenCalled();
    });
  });

  describe('buy', () => {
    it('should throw BadRequestException if userId is not the same as reqUserId', async () => {
      const params = { ticketId: 1, userId: 2, reqUserId: 3 };

      await expect(service.buy(params)).rejects.toThrow(BadRequestException);
      expect(repository.findOneBy).not.toHaveBeenCalled();
      expect(ordersService.create).not.toHaveBeenCalled();
      expect(processOrdersQueue.add).not.toHaveBeenCalled();
    });

    it('should throw NotFoundException if ticket is not found', async () => {
      const params = { ticketId: 1, userId: 2, reqUserId: 2 };
      repository.findOneBy.mockResolvedValue(null);

      await expect(service.buy(params)).rejects.toThrow(NotFoundException);
      expect(repository.findOneBy).toHaveBeenCalledWith({
        id: params.ticketId,
      });
      expect(ordersService.create).not.toHaveBeenCalled();
      expect(processOrdersQueue.add).not.toHaveBeenCalled();
    });

    it('should throw InternalServerErrorException if no tickets are available', async () => {
      const params = { ticketId: 1, userId: 2, reqUserId: 2 };
      repository.findOneBy.mockResolvedValue({
        id: 1,
        availableUnits: 0,
      } as Ticket);

      await expect(service.buy(params)).rejects.toThrow(
        InternalServerErrorException,
      );
      expect(repository.findOneBy).toHaveBeenCalledWith({
        id: params.ticketId,
      });
      expect(ordersService.create).not.toHaveBeenCalled();
      expect(processOrdersQueue.add).not.toHaveBeenCalled();
    });
  });
});
