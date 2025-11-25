package postusecase

import "github.com/bishal05das/blog_app_1/internal/domain"

type UpdatePostUseCase struct {
	repo domain.PostRepository
}

func NewUpdatePostUseCase(r domain.PostRepository) *UpdatePostUseCase {
	return &UpdatePostUseCase{
		repo: r,
	}
}

func (r *UpdatePostUseCase) Execute(p domain.Post) (*domain.Post,error) {
	return r.repo.UpdatePost(p)
}