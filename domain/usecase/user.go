package usecase

import (
	"freddy.facuellarg.com/domain/entities"
	"github.com/ansel1/merry"
	"gorm.io/gorm"
)

//UserRepositoryI public interface for user repository
type UserRepositoryI interface {
	CreateUser(user entities.User) (*entities.User, error)
	DeleteUser(id int) (entities.User, error)
	SearchUserByEmail(email string) (*entities.User, error)
}

//userRepository this contains all functionalities
//required for operate with user databases
type userRepository struct {
	db *gorm.DB
}

//NewUserRepository return a new userRepository
func NewUserRepository(db *gorm.DB) UserRepositoryI {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(user entities.User) (*entities.User, error) {
	result := ur.db.Create(&user)
	if result.Error != nil {
		return nil, merry.Wrap(result.Error)
	}
	return &user, nil

}

func (ur *userRepository) DeleteUser(id int) (entities.User, error) {
	panic("not implemented") // TODO: Implement
}

func (ur *userRepository) SearchUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := ur.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, merry.Wrap(result.Error)
	}
	return &user, nil
	//panic("not implemented") // TODO: Implement
}
