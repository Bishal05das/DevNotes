package reportusecase

import (
	"fmt"

	"github.com/bishal05das/blog_app_1/internal/domain"
)

type PostReportUseCase struct {
	repo domain.PostRepository
}

func NewPostReportUseCase(repo domain.PostRepository) *PostReportUseCase {
	return &PostReportUseCase{
		repo: repo,
	}
}

func (uc *PostReportUseCase) Generate(id int, formatter Formatter) string {
	post, err := uc.repo.GetPost(id)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	result := formatter.Format(post)
	return result
}
