package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rugi123/todo-api/internal/config"
	"github.com/rugi123/todo-api/internal/models"
	"github.com/rugi123/todo-api/internal/service"
)

type AuthHandler struct {
	Config  config.Config
	Service *service.Service
}

type Claims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
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
	}
	baseGroup := router.Group("/")
	{
		baseGroup.GET("/", h.ShowIndexPage)
		baseGroup.GET("/profile", AuthMiddleware(h.Config.AppConfig.JWTSecret), h.ShowProfilePage)
	}
}

func AuthMiddleware(jwtKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization required"})
				return
			}
			tokenString = authHeader[len("Bearer "):]
		}

		claims := &Claims{}
		jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

//Handlers

func (h *AuthHandler) ShowIndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "фронтенд кал",
	})
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var user models.User
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

	if err := service.CheckHashPassword(user.PasswordHash, credentials.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "password is incorect " + err.Error(),
		})
		return
	}

	expirationTime := time.Now().Add(5 + time.Minute)
	claims := &Claims{
		UserName: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.Config.AppConfig.JWTSecret))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token " + err.Error(),
		})
	}

	ctx.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (h *AuthHandler) ShowLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func (h *AuthHandler) ShowProfilePage(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*Claims)
	ctx.HTML(http.StatusOK, "profile.html", gin.H{
		"username": claims.UserName,
	})
}
