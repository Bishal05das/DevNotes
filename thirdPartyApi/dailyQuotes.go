package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func dailyNews(w http.ResponseWriter, r *http.Request) {
	news := []string{
		"daily news 1", "", "daily news 3", "",
	}
	rand.Seed(time.Now().UnixNano())
	SendData(w, news[rand.Intn(len(news))], http.StatusCreated)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", dailyNews)
	fmt.Println("server starting at port 3001")
	http.ListenAndServe(":3001", mux)
}

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}

}