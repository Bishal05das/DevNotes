package userusecase

import "github.com/bishal05das/blog_app_1/internal/domain"

type GetUserUseCase struct {
	repo domain.UserRepository
}

func NewGetUserUseCase(repo domain.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{
		repo: repo,
	}
}

func (uc *GetUserUseCase) Execute(id int) (*domain.User,error) {
	return uc.repo.GetUser(id)
}