package main

import (
	"fmt"
	"net/http"

	"github.com/bishal05das/blog_app_1/config"
	handler "github.com/bishal05das/blog_app_1/internal/delivery/handler"
	"github.com/bishal05das/blog_app_1/internal/repository"
	dailyquotesusecase "github.com/bishal05das/blog_app_1/internal/usecase/DailyQuotes"
	postusecase "github.com/bishal05das/blog_app_1/internal/usecase/post"
	userusecase "github.com/bishal05das/blog_app_1/internal/usecase/user"
	"github.com/bishal05das/blog_app_1/pkg/db"
)

func main() {
	mux := http.NewServeMux()

	cfg := config.GetConfig()

	db, err := db.NewConnection(cfg)
	if err != nil {
		fmt.Println("err in database connection: ", err)
		return
	}

	postRepo := repository.NewPostRepositoryDB(db)
	postHandler := handler.NewPostHandler(postusecase.NewCreatePostUseCase(postRepo),
		postusecase.NewGetPostByIDCase(postRepo),
		postusecase.NewListPostUseCase(postRepo))

	userRepo := repository.NewUserRepositoryDB(db)
	userHandler := handler.NewUserHandler(userusecase.NewCreateUserUseCase(userRepo),
		userusecase.NewGetUserUseCase(userRepo),
		userusecase.NewListUserUseCase(userRepo),
		userusecase.NewUpdateUserUseCase(userRepo))

	supportHandler := handler.NewSupportFactory(userRepo,postRepo)
	quotesHandler := handler.NewQuotesHnadler(dailyquotesusecase.NewDailyQuotes())

	mux.HandleFunc("POST /posts", postHandler.Create)
	mux.HandleFunc("GET /posts", postHandler.List)
	mux.HandleFunc("POST /users",userHandler.Create)
	mux.HandleFunc("GET /users",userHandler.List)
	mux.HandleFunc("POST /posts/{id}/support",supportHandler.Support)
	mux.HandleFunc("GET /dailyquotes",quotesHandler.GetQuotes)

	fmt.Println("Listening to the Server: 3000")
	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
