package supportusecase

import (
	"fmt"

	"github.com/bishal05das/blog_app_1/internal/domain"
)

type AdminSupportUseCase struct {
	UserRepo domain.UserRepository
	PostRepo domain.PostRepository
}

func NewAdminSupportUseCase(ur domain.UserRepository, pr domain.PostRepository) *AdminSupportUseCase {
	return &AdminSupportUseCase{UserRepo: ur, PostRepo: pr}
}

func (uc *AdminSupportUseCase) Execute(postID int) error {
	authorID, err := uc.PostRepo.GetAuthorIDByPostID(postID)
	if err != nil {
		fmt.Println("Invalid Post ID: ",err)
		return err
	}
	//fmt.Println(*authorID)
	return uc.UserRepo.UpdateReputation(*authorID)
}
