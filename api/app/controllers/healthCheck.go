package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// responseBodyで処理の成功を返す
	responseBody, err := getSuccessResponse()
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
