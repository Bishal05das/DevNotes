package userusecase

import "github.com/bishal05das/blog_app_1/internal/domain"

type UpdateUserUseCase struct {
	repo domain.UserRepository
}

func NewUpdateUserUseCase(repo domain.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		repo: repo,
	}
}

func (uc UpdateUserUseCase) Execute(user *domain.User) (*domain.User,error) {
	return uc.repo.UpdateUser(user)
}