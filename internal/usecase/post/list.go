package postusecase

import "github.com/bishal05das/blog_app_1/internal/domain"

type ListPostUseCase struct {
	repo domain.PostRepository
}

func NewListPostUseCase(r domain.PostRepository) *ListPostUseCase {
	return &ListPostUseCase{
		repo: r,
	}
}

func (r *ListPostUseCase) Execute() ([]*domain.Post, error) {
	return r.repo.ListPost()
}