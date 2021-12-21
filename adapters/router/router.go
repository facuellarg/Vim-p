package router

import (
	"freddy.facuellarg.com/domain/usecase"
	"github.com/labstack/echo/v4"
)

//router manage the routing process
type router struct {
	server   *echo.Echo
	userRepo usecase.UserRepositoryI
}

//NewRouter return a new router instance
func NewRouter(server *echo.Echo, userRepo usecase.UserRepositoryI) *router {
	return &router{
		server:   server,
		userRepo: userRepo,
	}
}

//RouteUsers put the routes for user
//operations in the server
func (r *router) RouteUsers() error {
	userRouter := NewUserRouter(r.userRepo)
	user := r.server.Group("/users")
	user.POST("", userRouter.RegistUser)
	return nil
}
