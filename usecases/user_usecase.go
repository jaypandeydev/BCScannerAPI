// usecases/user_usecase.go
package usecases

import (
	"BCScanner/domain"
	"BCScanner/interfaces"
)

type UserUsecase struct {
	UserRepository interfaces.UserRepository
}

func NewUserUsecase(repo interfaces.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepository: repo}
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.UserRepository.CreateUser(user)
}

func (u *UserUsecase) GetUserByID(id string) (*domain.User, error) {
	return u.UserRepository.GetUserByID(id)
}

func (u *UserUsecase) UpdateUser(user *domain.User) error {
	return u.UserRepository.UpdateUser(user)
}

func (u *UserUsecase) DeleteUser(id string) error {
	return u.UserRepository.DeleteUser(id)
}

func (u *UserUsecase) GetAllUsers() ([]domain.User, error) {
	return u.UserRepository.GetAllUsers()
}
