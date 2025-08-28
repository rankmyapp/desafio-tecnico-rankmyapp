package service

import (
    "strconv"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "desafio-tecnico/internal/config"
)

type JWTService interface {
    GenerateToken(userID uint) (string, error)
    ParseSubject(tokenStr string) (uint, error)
}

type jwtService struct {
    secret     []byte
    expiresIn  time.Duration
}

func NewJWTService(cfg *config.Config) JWTService {
    return &jwtService{
        secret:    []byte(cfg.JWTSecret),
        expiresIn: cfg.JWTExpiresIn,
    }
}

func (s *jwtService) GenerateToken(userID uint) (string, error) {
    now := time.Now().UTC()
    claims := jwt.RegisteredClaims{
        Subject:   strconv.Itoa(int(userID)),
        IssuedAt:  jwt.NewNumericDate(now),
        ExpiresAt: jwt.NewNumericDate(now.Add(s.expiresIn)),
    }
    tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return tkn.SignedString(s.secret)
}

func (s *jwtService) ParseSubject(tokenStr string) (uint, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
        return s.secret, nil
    })
    if err != nil {
        return 0, err
    }
    if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
        id64, err := strconv.ParseUint(claims.Subject, 10, 64)
        if err != nil {
            return 0, err
        }
        return uint(id64), nil
    }
    return 0, jwt.ErrTokenInvalidClaims
}