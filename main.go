package main

import (
	"github.com/dreisel/steroids-auth/auth"
	orm "github.com/dreisel/steroids-auth/db"
	"github.com/dreisel/steroids-auth/server"
	"github.com/dreisel/steroids-auth/users"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "auth ", log.LstdFlags|log.Lshortfile)
	server := server.New()
	db := orm.Connect()
	defer db.Close()
	userService := users.NewUserService(db)
	auth.SetRoutes(server, logger, userService)
	server.Run()
}
