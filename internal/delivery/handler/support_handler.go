package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bishal05das/blog_app_1/internal/domain"
	supportusecase "github.com/bishal05das/blog_app_1/internal/usecase/support"
	util "github.com/bishal05das/blog_app_1/utils"
)

type SupportFactory struct {
	UserRepo domain.UserRepository
	PostRepo domain.PostRepository
}

func NewSupportFactory(UserRepo domain.UserRepository, PostRepo domain.PostRepository) *SupportFactory {
	return &SupportFactory{
		UserRepo: UserRepo,
		PostRepo: PostRepo,
	}
}

func (h *SupportFactory) Support(w http.ResponseWriter, r *http.Request) {
	postidstr := r.PathValue("id")
	postID, err := strconv.Atoi(postidstr)
	if err != nil {
		fmt.Println("Error in converting id into int", err)
		return
	}

	usrRole := r.Header.Get("userRole")
	if usrRole == "admin" {
		supportusecase.NewAdminSupportUseCase(h.UserRepo,h.PostRepo).Execute(postID)
	}else {
	supportusecase.NewNormalSupportUseCase(h.PostRepo).Execute(postID)
	}
	util.SendData(w,"Successfully Supported",http.StatusOK)

}
