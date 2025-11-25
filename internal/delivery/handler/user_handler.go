package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bishal05das/blog_app_1/internal/domain"
	userusecase "github.com/bishal05das/blog_app_1/internal/usecase/user"
	util "github.com/bishal05das/blog_app_1/utils"
)

type UserHandler struct {
	createUC *userusecase.CreateUserUseCase
	getUC *userusecase.GetUserUseCase
	listUC *userusecase.ListUserUseCase
	updateUC *userusecase.UpdateUserUseCase
}

func NewUserHandler(createUC *userusecase.CreateUserUseCase,
	getUC *userusecase.GetUserUseCase,
	listUC *userusecase.ListUserUseCase,
	updateUC *userusecase.UpdateUserUseCase) *UserHandler {
		return &UserHandler{
			createUC: createUC,
			getUC: getUC,
			listUC: listUC,
			updateUC: updateUC,
		}
	}


func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	h.createUC.Execute(&user)
	util.SendData(w,user,http.StatusCreated)
	
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	var id int
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	user, err :=h.getUC.Execute(id)
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	util.SendData(w,user,http.StatusFound)
}

func (h *UserHandler) List(w http.ResponseWriter,r *http.Request) {
	users, err := h.listUC.Execute()
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	util.SendData(w,users,http.StatusFound)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	h.updateUC.Execute(&user)
	util.SendData(w,user,http.StatusOK)

}