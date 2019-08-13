package auth

import (
	"github.com/dreisel/steroids-auth/users"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Routes struct {
	logger       *log.Logger
	usersService users.UserService
	ac           authCookie
}

type LoginRequest struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func (r Routes) Login(c *gin.Context) {
	var json LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := r.usersService.GetByUsernameAndPassword(json.Username, json.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
		return
	}

	err = r.ac.set(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	r.logger.Printf("USER %d LOGGED IN SUCCESSFULLY", user.ID)
	c.JSON(http.StatusCreated, gin.H{"username": user.Username, "id": user.ID})
}

func (r Routes) DeRegister(c *gin.Context) {
	claims, err := r.ac.get(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token", "message": err.Error()})
		return
	}
	uid, err := strconv.Atoi(claims.Subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid subject token", "uid": uid})
		return
	}
	user, err := r.usersService.Delete(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": user.ID})
}

func (r Routes) Logout(c *gin.Context) {
	r.ac.delete(c)
	c.JSON(http.StatusOK, gin.H{})
}

func (r Routes) Register(c *gin.Context) {
	var json LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := r.usersService.Create(json.Username, json.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = r.ac.set(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"username": user.Username, "id": user.ID})
}

func newRoutes(logger *log.Logger, usersService users.UserService) Routes {
	routes := Routes{
		logger:       logger,
		usersService: usersService,
	}
	return routes
}

func SetRoutes(server *gin.Engine, logger *log.Logger, usersService users.UserService) {
	routes := newRoutes(logger, usersService)
	server.POST("/login", routes.Login)
	server.POST("/register", routes.Register)
	server.GET("/logout", routes.Logout)
	server.GET("/deregister", routes.DeRegister)
}
