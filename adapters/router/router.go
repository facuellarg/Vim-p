package router

import (
	"freddy.facuellarg.com/domain/usecase"
	"github.com/labstack/echo/v4"
)

//router manage the routing process
type router struct {
	userRepo usecase.UserRepositoryI
}

//NewRouter return a new router instance
func NewRouter(userRepo usecase.UserRepositoryI) *router {
	return &router{
		userRepo: userRepo,
	}
}

//RouteUsers put the routes for user
//operations in the server
func (r *router) RouteUsers(server *echo.Echo) error {
	userRouter := NewUserRouter(r.userRepo)
	user := server.Group("/users")
	user.POST("", userRouter.RegistUser)
	user.POST("/loggin", userRouter.Loggin)
	return nil
}
