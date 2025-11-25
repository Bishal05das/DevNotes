package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bishal05das/blog_app_1/internal/domain"
	postusecase "github.com/bishal05das/blog_app_1/internal/usecase/post"
	util "github.com/bishal05das/blog_app_1/utils"
)

type PostHandler struct {
	createUC *postusecase.CreatePostUseCase
	getUC *postusecase.GetPostByIDUseCase
	listUC *postusecase.ListPostUseCase
}

func NewPostHandler(createUC *postusecase.CreatePostUseCase,
	getUC *postusecase.GetPostByIDUseCase,listUC *postusecase.ListPostUseCase) *PostHandler {
	return &PostHandler{
		createUC: createUC,
		getUC: getUC,
		listUC: listUC,
	}
}

func (h *PostHandler) Create(w http.ResponseWriter,r *http.Request) {
	var post domain.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	h.createUC.Execute(&post)
	util.SendData(w,post,http.StatusCreated)
}

func (h *PostHandler) Get(w http.ResponseWriter,r *http.Request) {
	var id int
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	post, err :=h.getUC.Execute(id)
	if err != nil {
		util.SendData(w,err,http.StatusBadRequest)
	}

	util.SendData(w,post,http.StatusFound)
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	posts,err := h.listUC.Execute()
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	util.SendData(w,posts,http.StatusOK)
}