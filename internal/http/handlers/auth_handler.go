package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "desafio-tecnico/internal/config"
    "desafio-tecnico/internal/service"
    "desafio-tecnico/pkg/queue"
)

type AuthHandler struct {
    cfg     *config.Config
    userSvc service.UserService
    pub     queue.Publisher
}

func NewAuthHandler(cfg *config.Config, userSvc service.UserService, pub queue.Publisher) *AuthHandler {
    return &AuthHandler{cfg: cfg, userSvc: userSvc, pub: pub}
}

type signupReq struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}
type loginReq struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *AuthHandler) Signup(c *gin.Context) {
    var req signupReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }
    user, token, err := h.userSvc.Signup(req.Name, req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Publica email de boas vindas
    payload := map[string]string{
        "to":      user.Email,
        "subject": "Bem-vindo!",
        "body":    fmt.Sprintf("Olá %s, bem-vindo à plataforma!", user.Name),
    }
    data, _ := json.Marshal(payload)
    _ = h.pub.Publish(h.cfg.EmailQueueName, data) // log/ignorar erro neste fluxo

    c.JSON(http.StatusCreated, gin.H{
        "user":  user,
        "token": token,
    })
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req loginReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }
    user, token, err := h.userSvc.Login(req.Email, req.Password)
    if err != nil {
        status := http.StatusUnauthorized
        if err.Error() == "invalid credentials" {
            status = http.StatusUnauthorized
        }
        c.JSON(status, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "user":  user,
        "token": token,
    })
}