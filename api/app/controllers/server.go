package controllers

import (
	"api/app/config"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type successResponse struct {
	Success bool `json:"success"`
}

// 処理の成功結果をレスポンスで返す
func getSuccessResponse() (responseBody []byte, err error) {
	responseBody, err = json.Marshal(successResponse{Success: true})
	return responseBody, err
}

func StartMainServer() error {
	// gorilla/muxを使ったルーティング
	r := mux.NewRouter().StrictSlash(true)
	port := fmt.Sprintf(":%s", config.Config.Port)

	// ヘルスチェック
	r.HandleFunc("/api/health_check", healthCheckHandler).Methods("GET")

	// User
	r.HandleFunc("/api/signup", signupHandler).Methods("POST")

	// corsの設定
	customizeCors := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
	c := customizeCors.Handler(r)
	return http.ListenAndServe(port, c)
}
