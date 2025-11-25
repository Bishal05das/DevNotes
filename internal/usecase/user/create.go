package userusecase

import "github.com/bishal05das/blog_app_1/internal/domain"

type CreateUserUseCase struct {
	repo domain.UserRepository
}

func NewCreateUserUseCase(userRepo domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: userRepo,
	}
}

func (uc *CreateUserUseCase) Execute(user *domain.User) error {
	return uc.repo.CreateUser(user)
}