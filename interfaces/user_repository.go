// interfaces/user_repository.go
package interfaces

import "BCScanner/domain"

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id string) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
	GetAllUsers() ([]domain.User, error)
}
