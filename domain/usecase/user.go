package usecase

import (
	"freddy.facuellarg.com/domain/entities"
	"github.com/ansel1/merry"
	"github.com/doug-martin/goqu/v9"
)

//UserRepositoryI public interface for user repository
type UserRepositoryI interface {
	CreateUser(user entities.User) (*entities.User, error)
	DeleteUser(id int) (entities.User, error)
	SearchUserByEmail(email string) (entities.User, error)
}

//userRepository this contains all functionalities
//required for operate with user databases
type userRepository struct {
	db *goqu.Database
}

//NewUserRepository return a new userRepository
func NewUserRepository(db *goqu.Database) UserRepositoryI {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(user entities.User) (*entities.User, error) {
	result, err := ur.db.Insert("users").Rows(
		user,
	).Executor().Exec()
	if err != nil {
		return nil, merry.Wrap(err)
	}
	user.Id, _ = result.LastInsertId()

	return &user, nil

}

func (ur *userRepository) DeleteUser(id int) (entities.User, error) {
	panic("not implemented") // TODO: Implement
}

func (ur *userRepository) SearchUserByEmail(email string) (entities.User, error) {
	panic("not implemented") // TODO: Implement
}
