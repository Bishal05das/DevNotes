package supportusecase

import (

	"github.com/bishal05das/blog_app_1/internal/domain"
)

type NormalSupportUseCase struct {
	repo domain.PostRepository
}

func NewNormalSupportUseCase(r domain.PostRepository) *NormalSupportUseCase {
	return &NormalSupportUseCase{repo: r}
}

func (uc *NormalSupportUseCase) Execute(postID int) error {
	//fmt.Println(postID)
	return uc.repo.UpdateScore(postID)
}
