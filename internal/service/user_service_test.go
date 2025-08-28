package service

import (
	"testing"

	"desafio-tecnico/internal/config"
)

func TestPasswordAndJWT(t *testing.T) {
	cfg := &config.Config{JWTSecret: "test", JWTExpiresIn: 0}
	jwtSvc := NewJWTService(cfg)

	us := &userService{
		repo:     nil,
		jwt:      jwtSvc,
		validate: nil,
	}

	// Hash e verificação
	pass := "StrongPass123!"
	u, token, err := us.Signup("Name", "email@example.com", pass)
	if err == nil && u != nil {
		t.Errorf("Signup should fail without repository, got user")
	}
	// Testa geração de token isoladamente
	tokenStr, err := jwtSvc.GenerateToken(42)
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}
	id, err := jwtSvc.ParseSubject(tokenStr)
	if err != nil {
		t.Fatalf("ParseSubject error: %v", err)
	}
	if id != 42 {
		t.Fatalf("expected id 42, got %d", id)
	}
	_ = token // keep refs
}
