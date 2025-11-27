package reportusecase

import (
	"fmt"

	"github.com/bishal05das/blog_app_1/internal/domain"
)

type UserReportUseCase struct {
	repo domain.UserRepository
}

func NewUserReportUseCase(repo domain.UserRepository) *UserReportUseCase {
	return &UserReportUseCase{
		repo: repo,
	}
}

func (uc *UserReportUseCase) Generate(id int, formatter Formatter) string {
	user, err := uc.repo.GetUser(id)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	result := formatter.Format(user)
	return result
}
