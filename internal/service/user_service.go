package service

import (
    "errors"
    "strings"

    "github.com/go-playground/validator/v10"
    "desafio-tecnico/internal/domain"
    "desafio-tecnico/internal/repository"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

var (
    ErrInvalidCredentials = errors.New("invalid credentials")
    ErrEmailInUse         = errors.New("email already in use")
)

type UserService interface {
    Signup(name, email, password string) (*domain.User, string, error)
    Login(email, password string) (*domain.User, string, error)
    GetByID(id uint) (*domain.User, error)
    List(search, sort string, page, limit int) ([]domain.User, int64, error)
    Update(id uint, name, email, password string) (*domain.User, error)
    Delete(id uint) error
}

type userService struct {
    repo    repository.UserRepository
    jwt     JWTService
    validate *validator.Validate
}

func NewUserService(repo repository.UserRepository, jwt JWTService) UserService {
    return &userService{
        repo: repo,
        jwt:  jwt,
        validate: validator.New(),
    }
}

func (s *userService) Signup(name, email, password string) (*domain.User, string, error) {
    email = strings.TrimSpace(strings.ToLower(email))
    if err := s.validate.Var(email, "required,email"); err != nil {
        return nil, "", err
    }
    if len(password) < 8 {
        return nil, "", errors.New("password must be at least 8 characters")
    }
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, "", err
    }
    u := &domain.User{
        Name:         strings.TrimSpace(name),
        Email:        email,
        PasswordHash: string(hash),
    }
    if err := s.repo.Create(u); err != nil {
        if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "duplicate key") {
            return nil, "", ErrEmailInUse
        }
        return nil, "", err
    }
    token, err := s.jwt.GenerateToken(u.ID)
    return u, token, err
}

func (s *userService) Login(email, password string) (*domain.User, string, error) {
    email = strings.TrimSpace(strings.ToLower(email))
    u, err := s.repo.GetByEmail(email)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, "", ErrInvalidCredentials
        }
        return nil, "", err
    }
    if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
        return nil, "", ErrInvalidCredentials
    }
    token, err := s.jwt.GenerateToken(u.ID)
    return u, token, err
}

func (s *userService) GetByID(id uint) (*domain.User, error) {
    return s.repo.GetByID(id)
}

func (s *userService) List(search, sort string, page, limit int) ([]domain.User, int64, error) {
    if page < 1 {
        page = 1
    }
    if limit <= 0 || limit > 100 {
        limit = 10
    }
    return s.repo.List(search, sort, page, limit)
}

func (s *userService) Update(id uint, name, email, password string) (*domain.User, error) {
    u, err := s.repo.GetByID(id)
    if err != nil {
        return nil, err
    }
    if name != "" {
        u.Name = strings.TrimSpace(name)
    }
    if email != "" {
        e := strings.TrimSpace(strings.ToLower(email))
        if err := s.validate.Var(e, "email"); err != nil {
            return nil, err
        }
        u.Email = e
    }
    if password != "" {
        if len(password) < 8 {
            return nil, errors.New("password must be at least 8 characters")
        }
        hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            return nil, err
        }
        u.PasswordHash = string(hash)
    }
    if err := s.repo.Update(u); err != nil {
        return nil, err
    }
    return u, nil
}

func (s *userService) Delete(id uint) error {
    return s.repo.Delete(id)
}