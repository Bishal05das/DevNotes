package postusecase

import "github.com/bishal05das/blog_app_1/internal/domain"

type CreatePostUseCase struct {
	repo domain.PostRepository
}

func NewCreatePostUseCase(r domain.PostRepository) *CreatePostUseCase {
	return &CreatePostUseCase{
		repo: r,
	}
}

func (uc *CreatePostUseCase) Execute(post *domain.Post) error {
	return uc.repo.CreatePost(post)
}