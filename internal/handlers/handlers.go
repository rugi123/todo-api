package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rugi123/todo-api/internal/config"
	"github.com/rugi123/todo-api/internal/models"
	"github.com/rugi123/todo-api/internal/service"
)

type AuthHandler struct {
	Config  config.Config
	Service *service.Service
}

func NewAuthHandler(cfg config.Config, service service.Service) *AuthHandler {
	return &AuthHandler{
		Config:  cfg,
		Service: &service,
	}
}

func (h *AuthHandler) RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.GET("/register", h.ShowRegiserPage)
		authGroup.GET("/login", h.ShowLoginPage)

		authGroup.POST("/register", h.Register)
		authGroup.POST("/login", h.Login)

		authGroup.GET("/profile")
	}
	baseGroup := router.Group("/")
	{
		baseGroup.GET("/", h.ShowIndexPage)
	}
}

func (h *AuthHandler) ShowIndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "фронтенд кал",
	})
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var user models.User
	fmt.Println(ctx.Params)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := h.Service.Save(ctx, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *AuthHandler) ShowRegiserPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", nil)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var credentials struct {
		Login    string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindBodyWithJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to should bind body" + err.Error(),
		})
		return
	}

	user, err := h.Service.Storage.UserStorage.GetUserByName(ctx, credentials.Login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}

	fmt.Println(service.CheckHashPassword("$2a$10$atQsVW1rTFl6in5eHxa30eFdZBaMSNi7jt2zFVrSIihCnbPOF0vNq", "1"))

	if err := service.CheckHashPassword(user.PasswordHash, credentials.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "password is incorect " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "все кайф",
	})
}

func (h *AuthHandler) ShowLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}
