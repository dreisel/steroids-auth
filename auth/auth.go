package auth

import (
	"github.com/dreisel/steroids-auth/users"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Routes struct {
	logger       *log.Logger
	usersService users.UserService
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
	user, err := r.usersService.Create(json.Username, json.Password)
	if  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"username": user.Username, "id": user.ID})
}

func (r Routes) Logout(c *gin.Context) {

}

func (r Routes) Register(c *gin.Context) {

}

func (r Routes) DeRegister(c *gin.Context) {

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
