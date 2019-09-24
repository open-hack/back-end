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

func (as *ApiServer) GetHackathonHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)

	hackathon, err := as.DB.GetHackathon(idInt)
	if err != nil {
		log.Println("error on get hackathon", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(hackathon); err != nil {
		log.Println("error encode hackathon object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) GetAllHackathonsHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	hackathons, err := as.DB.GetAllHackathons()
	if err != nil {
		log.Println("error on get hackathons", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(hackathons); err != nil {
		log.Println("error encode hackathon object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) CreateHackathonHandle(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	//Leitura do body da requisição
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Error on getting body content", err)
		http.Error(w, "Error on getting body content", 500)
		return
	}
	hackathon := &model.Hackathon{}
	err = json.Unmarshal(b, &hackathon)
	if err != nil {
		log.Println("Error on unmarshal info from body", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}

	if err := as.DB.CreateHackathon(hackathon); err != nil {
		log.Println("error on create hackathon data", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}
}
