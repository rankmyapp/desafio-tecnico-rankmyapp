import { Test, TestingModule } from '@nestjs/testing';
import { UsersService } from './users.service';
import { getRepositoryToken } from '@nestjs/typeorm';
import { User } from './users.entity';
import { Repository } from 'typeorm';

describe('UsersService', () => {
  let service: UsersService;
  let repository: jest.Mocked<Repository<User>>;

  const mockRepository = () => ({
    create: jest.fn(),
    save: jest.fn(),
    findOneBy: jest.fn(),
    find: jest.fn(),
    remove: jest.fn(),
  });

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        UsersService,
        {
          provide: getRepositoryToken(User),
          useFactory: mockRepository,
        },
      ],
    }).compile();

    service = module.get<UsersService>(UsersService);
    repository = module.get(getRepositoryToken(User));
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  describe('create', () => {
    it('should create and save a new user', async () => {
      const userData = {
        email: 'test@example.com',
        password: 'password123',
      };
      const createdUser = { ...userData, id: 1 } as User;

      repository.create.mockReturnValue(createdUser);
      repository.save.mockResolvedValue(createdUser);

      const result = await service.create(userData);

      expect(repository.create).toHaveBeenCalledWith(userData);
      expect(repository.save).toHaveBeenCalledWith(createdUser);
      expect(result).toEqual(createdUser);
    });
  });

  describe('findOne', () => {
    it('should return a user if found', async () => {
      const mockUser = {
        id: 1,
        email: 'test@example.com',
        password: 'password123',
      } as User;
      repository.findOneBy.mockResolvedValue(mockUser);

      const result = await service.findOne(1);

      expect(repository.findOneBy).toHaveBeenCalledWith({ id: 1 });
      expect(result).toEqual(mockUser);
    });

    it('should return null if user is not found', async () => {
      repository.findOneBy.mockResolvedValue(null);

      const result = await service.findOne(999);

      expect(repository.findOneBy).toHaveBeenCalledWith({ id: 999 });
      expect(result).toBeNull();
    });
  });

  describe('find', () => {
    it('should return users with matching email', async () => {
      const email = 'test@example.com';
      const mockUsers = [{ id: 1, email, password: 'password123' }] as User[];
      repository.find.mockResolvedValue(mockUsers);

      const result = await service.find(email);

      expect(repository.find).toHaveBeenCalledWith({ where: { email } });
      expect(result).toEqual(mockUsers);
    });

    it('should return empty array if no users match the email', async () => {
      const email = 'nonexistent@example.com';
      repository.find.mockResolvedValue([]);

      const result = await service.find(email);

      expect(repository.find).toHaveBeenCalledWith({ where: { email } });
      expect(result).toEqual([]);
    });
  });

  describe('update', () => {
    it('should update a user if found', async () => {
      const userId = 1;
      const updateData = { email: 'updated@example.com' };
      const existingUser = {
        id: userId,
        email: 'test@example.com',
        password: 'password123',
      } as User;
      const updatedUser = { ...existingUser, ...updateData } as User;

      repository.findOneBy.mockResolvedValue(existingUser);
      repository.save.mockResolvedValue(updatedUser);

      const result = await service.update(userId, updateData);

      expect(repository.findOneBy).toHaveBeenCalledWith({ id: userId });
      expect(repository.save).toHaveBeenCalledWith(updatedUser);
      expect(result).toEqual(updatedUser);
    });

    it('should throw Error if user is not found', async () => {
      repository.findOneBy.mockResolvedValue(null);

      await expect(
        service.update(999, { email: 'updated@example.com' }),
      ).rejects.toThrow('User not found');
      expect(repository.findOneBy).toHaveBeenCalledWith({ id: 999 });
      expect(repository.save).not.toHaveBeenCalled();
    });
  });

  describe('remove', () => {
    it('should remove a user if found', async () => {
      const userId = 1;
      const existingUser = {
        id: userId,
        email: 'test@example.com',
        password: 'password123',
      } as User;

      repository.findOneBy.mockResolvedValue(existingUser);
      repository.remove.mockResolvedValue(existingUser);

      const result = await service.remove(userId);

      expect(repository.findOneBy).toHaveBeenCalledWith({ id: userId });
      expect(repository.remove).toHaveBeenCalledWith(existingUser);
      expect(result).toEqual(existingUser);
    });

    it('should throw Error if user is not found', async () => {
      repository.findOneBy.mockResolvedValue(null);

      await expect(service.remove(999)).rejects.toThrow('User not found');
      expect(repository.findOneBy).toHaveBeenCalledWith({ id: 999 });
      expect(repository.remove).not.toHaveBeenCalled();
    });
  });
});
