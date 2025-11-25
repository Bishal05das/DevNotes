package dailyquotesusecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DailyQuotesUseCase struct{}

func NewDailyQuotes() *DailyQuotesUseCase {
	return &DailyQuotesUseCase{}
}

func (dq *DailyQuotesUseCase) Execute() (*string,error) {
	for i := 0; i < 5; i++ {
		resp, err := http.Get("http://localhost:3001")

		if err != nil {
			fmt.Println("error fetching daily news", err)
			return nil,err
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			return nil,err
		}
		resp.Body.Close()
		var quotes string
		if err := json.Unmarshal(body, &quotes); err != nil {
			fmt.Println("error unmarshaling response", err)
			continue
		}
		if quotes != "" {
			return &quotes,nil
		}
		time.Sleep(1 * time.Second)
	}
	return nil,fmt.Errorf("can not get any quotes")
}
