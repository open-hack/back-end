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

func (as *ApiServer) GetSubscriptionHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)

	subscription, err := as.DB.GetSubscription(idInt)
	if err != nil {
		log.Println("error on get subscription", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(subscription); err != nil {
		log.Println("error encode subscription object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) GetAllSubscriptionsHandle(w http.ResponseWriter, r *http.Request) {

	subscriptions, err := as.DB.GetAllSubscriptions()
	if err != nil {
		log.Println("error on get subscriptions", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(subscriptions); err != nil {
		log.Println("error encode subscription object", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (as *ApiServer) CreateSubscriptionHandle(w http.ResponseWriter, r *http.Request) {
	//Leitura do body da requisição
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Error on getting body content", err)
		http.Error(w, "Error on getting body content", 500)
		return
	}
	subscription := &model.Subscription{}
	err = json.Unmarshal(b, &subscription)
	if err != nil {
		log.Println("Error on unmarshal info from body", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}

	if err := as.DB.CreateSubscription(subscription); err != nil {
		log.Println("error on create subscription data", err)
		http.Error(w, "Error on unmarshal info from body", 500)
		return
	}
}
