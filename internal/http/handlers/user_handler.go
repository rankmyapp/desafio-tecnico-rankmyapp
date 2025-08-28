package handlers

import (
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
    "desafio-tecnico/internal/http/middleware"
    "desafio-tecnico/internal/service"
)

const (INVALID_ID = "invalid id")

type UserHandler struct {
    userSvc service.UserService
}

func NewUserHandler(userSvc service.UserService) *UserHandler {
    return &UserHandler{userSvc: userSvc}
}

func (h *UserHandler) List(c *gin.Context) {
    search := strings.TrimSpace(c.Query("search"))
    sort := strings.TrimSpace(c.Query("sort"))
    page := atoiDefault(c.Query("page"), 1)
    limit := atoiDefault(c.Query("limit"), 10)

    users, total, err := h.userSvc.List(search, sort, page, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "items": users,
        "meta": gin.H{
            "total": total,
            "page":  page,
            "limit": limit,
        },
    })
}

func (h *UserHandler) GetByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": INVALID_ID})
        return
    }
    u, err := h.userSvc.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    c.JSON(http.StatusOK, u)
}

type updateReq struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *UserHandler) Update(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": INVALID_ID})
        return
    }
    var req updateReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }

    // Opcional: só permitir update do próprio usuário
    if uid, ok := c.Get(middleware.CtxUserID); ok {
        if uint(id) != uid.(uint) {
            c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
            return
        }
    }

    u, err := h.userSvc.Update(uint(id), req.Name, req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, u)
}

func (h *UserHandler) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": INVALID_ID})
        return
    }

    if uid, ok := c.Get(middleware.CtxUserID); ok {
        if uint(id) != uid.(uint) {
            c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
            return
        }
    }
    if err := h.userSvc.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete"})
        return
    }
    c.Status(http.StatusNoContent)
}

func atoiDefault(s string, def int) int {
    if s == "" {
        return def
    }
    i, err := strconv.Atoi(s)
    if err != nil {
        return def
    }
    return i
}