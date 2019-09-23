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

func (as *ApiServer) GetHackathonUserHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)

	hackathonUser, err := as.DB.GetHackathonUser(idInt)
	if err != nil {
		log.Println("error on get hackathonUser", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(hackathonUser); err != nil {
		log.Println("error encode hackathonUser object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) GetAllHackathonUsersHandle(w http.ResponseWriter, r *http.Request) {

	hackathonUsers, err := as.DB.GetAllHackathonUsers()
	if err != nil {
		log.Println("error on get hackathonUsers", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(hackathonUsers); err != nil {
		log.Println("error encode hackathonUser object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) CreateHackathonUserHandle(w http.ResponseWriter, r *http.Request) {
	//Leitura do body da requisição
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Error on getting body content", err)
		http.Error(w, "Error on getting body content", 500)
		return
	}
	hackathonUser := &model.HackathonUser{}
	err = json.Unmarshal(b, &hackathonUser)
	if err != nil {
		log.Println("Error on unmarshal info from body", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}

	if err := as.DB.CreateHackathonUser(hackathonUser); err != nil {
		log.Println("error on create hackathonUser data", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}
}

func (as *ApiServer) CreateByUserIDHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userID, _ := strconv.ParseInt(id, 10, 64)

	//Leitura do body da requisição
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Error on getting body content", err)
		http.Error(w, "Error on getting body content", 500)
		return
	}

	var hackathonIDs []int64
	err = json.Unmarshal(b, &hackathonIDs)
	if err != nil {
		log.Println("Error on unmarshal info from body", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}

	if err := as.DB.CreateByUserID(userID, hackathonIDs); err != nil {
		log.Println("error on create hackathonUser data", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}
}

func (as *ApiServer) CreateByHackathonIDHandle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	hackathonID, _ := strconv.ParseInt(id, 10, 64)

	//Leitura do body da requisição
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Error on getting body content", err)
		http.Error(w, "Error on getting body content", 500)
		return
	}

	var userIDs []int64
	err = json.Unmarshal(b, &userIDs)
	if err != nil {
		log.Println("Error on unmarshal info from body", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}

	if err := as.DB.CreateByHackathonID(hackathonID, userIDs); err != nil {
		log.Println("error on create hackathonUser data", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}
}
