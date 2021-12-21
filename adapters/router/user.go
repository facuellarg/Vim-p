package router

import (
	"fmt"
	"net/http"

	"freddy.facuellarg.com/domain/entities"
	"freddy.facuellarg.com/domain/usecase"
	"github.com/ansel1/merry"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

//userRoutes implements routes
//for user
type userRoutes struct {
	userResitory usecase.UserRepositoryI
}

//NewUserRouter return a new user router
func NewUserRouter(repository usecase.UserRepositoryI) *userRoutes {
	return &userRoutes{
		userResitory: repository,
	}
}

//RegisterUser create a user
func (ur *userRoutes) RegistUser(c echo.Context) error {
	var user entities.User
	if err := c.Bind(&user); err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	fmt.Printf("%+v\n", user)
	//user.Birthdate = entities.CustomTime(time.Now())
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return merry.Wrap(err)
	}
	user.Password = string(hashedPass)
	userCreated, err := ur.userResitory.CreateUser(user)
	if err != nil {
		return merry.Wrap(err)
	}
	return c.JSON(http.StatusAccepted, userCreated)
}
