package userusecase


import "github.com/bishal05das/blog_app_1/internal/domain"

type ListUserUseCase struct {
	repo domain.UserRepository
}

func NewListUserUseCase(repo domain.UserRepository) *ListUserUseCase {
	return &ListUserUseCase{
		repo: repo,
	}
}

func (uc *ListUserUseCase) Execute() ([]*domain.User,error) {
	return uc.repo.ListUser()
}