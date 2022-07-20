package controllers

import (
	"api/app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	signUpUser := &models.User{}
	if err := json.Unmarshal(req, &signUpUser); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	validate := validator.New()
	if err := validate.Struct(signUpUser); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
		return
	}

	if err := signUpUser.CreateUser(); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	res, err := getSuccessResponse()
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
