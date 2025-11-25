package postusecase

import "github.com/bishal05das/blog_app_1/internal/domain"

type GetPostByIDUseCase struct {
	repo domain.PostRepository
}

func NewGetPostByIDCase(r domain.PostRepository) *GetPostByIDUseCase {
	return &GetPostByIDUseCase{
		repo: r,
	}
}

func (uc *GetPostByIDUseCase) Execute(id int) (*domain.Post,error) {
	return uc.repo.GetPost(id)
}