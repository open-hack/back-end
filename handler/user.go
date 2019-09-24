package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/open-hack/back-end/model"
)

func (as *ApiServer) GetUserHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)

	user, err := as.DB.GetUser(idInt)
	if err != nil {
		log.Println("error on get user", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("error encode user object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) GetAllUsersHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	users, err := as.DB.GetAllUsers()
	if err != nil {
		log.Println("error on get users", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Println("error encode user object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

// user login
func (as *ApiServer) LoginUserHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	vars := mux.Vars(r)
	email := vars["email"]

	user, err := as.DB.GetUserByEmail(email)
	if err != nil {
		log.Println("error on get user", err)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userUpdated, error := as.DB.UpdateLastLogin(user)
	if error != nil {
		log.Println("error on get user", error)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if encodeErr := json.NewEncoder(w).Encode(userUpdated); encodeErr != nil {
		log.Println("error encode user object", encodeErr)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) CreateUserHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	//Leitura do body da requisição
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Error on getting body content", err)
		http.Error(w, "Error on getting body content", 500)
		return
	}
	user := &model.User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Println("Error on unmarshal info from body", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}

	if err := as.DB.CreateUser(user); err != nil {
		log.Println("error on create user data", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}
}
