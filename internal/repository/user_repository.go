package repository

import (
    "strings"

    "desafio-tecnico/internal/domain"
    "gorm.io/gorm"
)

type UserRepository interface {
    Create(u *domain.User) error
    GetByEmail(email string) (*domain.User, error)
    GetByID(id uint) (*domain.User, error)
    List(search, sort string, page, limit int) ([]domain.User, int64, error)
    Update(u *domain.User) error
    Delete(id uint) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(u *domain.User) error {
    return r.db.Create(u).Error
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
    var u domain.User
    if err := r.db.Where("LOWER(email) = ?", strings.ToLower(email)).First(&u).Error; err != nil {
        return nil, err
    }
    return &u, nil
}

func (r *userRepository) GetByID(id uint) (*domain.User, error) {
    var u domain.User
    if err := r.db.First(&u, id).Error; err != nil {
        return nil, err
    }
    return &u, nil
}

func (r *userRepository) List(search, sort string, page, limit int) ([]domain.User, int64, error) {
    var users []domain.User
    q := r.db.Model(&domain.User{})

    if search != "" {
        s := "%" + strings.ToLower(search) + "%"
        q = q.Where("LOWER(name) LIKE ? OR LOWER(email) LIKE ?", s, s)
    }

    var total int64
    if err := q.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    switch sort {
    case "name asc":
        q = q.Order("name asc")
    case "name desc":
        q = q.Order("name desc")
    case "created_at asc":
        q = q.Order("created_at asc")
    default:
        q = q.Order("created_at desc")
    }

    offset := (page - 1) * limit
    if err := q.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
        return nil, 0, err
    }

    return users, total, nil
}

func (r *userRepository) Update(u *domain.User) error {
    return r.db.Save(u).Error
}

func (r *userRepository) Delete(id uint) error {
    return r.db.Delete(&domain.User{}, id).Error
}