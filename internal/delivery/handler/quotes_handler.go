package handler

import (
	"net/http"

	dailyquotesusecase "github.com/bishal05das/blog_app_1/internal/usecase/DailyQuotes"
	util "github.com/bishal05das/blog_app_1/utils"
)

type QuotesHandler struct {
	quotesUC *dailyquotesusecase.DailyQuotesUseCase
}

func NewQuotesHnadler(quotesUC *dailyquotesusecase.DailyQuotesUseCase) *QuotesHandler {
	return &QuotesHandler{
		quotesUC: quotesUC,
	}
}

func (qh *QuotesHandler) GetQuotes(w http.ResponseWriter, r *http.Request){
	quotes,err := qh.quotesUC.Execute()
	if err != nil {
		util.SendData(w,"error getting quotes",http.StatusBadRequest)
	}
	util.SendData(w,quotes,http.StatusOK)
}