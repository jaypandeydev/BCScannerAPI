// usecases/user_usecase.go
package usecases

import (
	"BCScanner/domain"
	"context"
	"time"
	//"BCScanner/repository"
)

type UserUsecase struct {
	UserRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (u *UserUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.UserRepository.Create(ctx, user)
}

func (u *UserUsecase) Fetch(c context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.UserRepository.Fetch(ctx)
}

func (u *UserUsecase) GetByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.UserRepository.GetByEmail(ctx, email)
}

func (u *UserUsecase) GetByID(c context.Context, Id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.UserRepository.GetByID(ctx, Id)
}

// func (u *UserUsecase) UpdateUser(c context.Context, user *domain.User) error {
// 	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
// 	defer cancel()
// 	return u.UserRepository.UpdateUser(user)
// }

// func (u *UserUsecase) DeleteUser(c context.Context, id string) error {
// 	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
// 	defer cancel()
// 	return u.UserRepository.DeleteUser(id)
// }
