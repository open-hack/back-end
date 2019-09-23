package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/open-hack/back-end/handler"
	"github.com/open-hack/back-end/model"
	"github.com/open-hack/back-end/service"
	"gopkg.in/olivere/elastic.v6"
)

func main() {

	routes := mux.NewRouter()

	db := model.WeeHackDB{}
	//Inicialização do repositorio
	db.MustInit()

	if err := godotenv.Load(); err != nil {
		log.Println("File .env not found, reading configuration from ENV")
	}

	elasticClient, err := elastic.NewClient(elastic.SetURL(model.ElasticSearch()), elastic.SetSniff(false))
	if err != nil {
		log.Fatal("Error Creating Elastic Client: ", err)
	}

	log.Printf("Elastic Search Client Created")

	elasticService := &service.ElasticService{
		ElasticCLI: elasticClient,
	}

	chatHandler := &handler.Handler{
		Upgrader: websocket.Upgrader{},
		Service:  elasticService,
	}

	apiServer := handler.ApiServer{
		DB: db,
	}

	//Rotas de consulta
	routes.HandleFunc("/api/hackathonUser/{id:[0-9]+}", apiServer.GetHackathonUserHandle).Methods("GET")
	routes.HandleFunc("/api/hackathonUser/all", apiServer.GetAllHackathonUsersHandle).Methods("GET")

	routes.HandleFunc("/api/hackathon/{id:[0-9]+}", apiServer.GetHackathonHandle).Methods("GET")
	routes.HandleFunc("/api/hackathon/all", apiServer.GetAllHackathonsHandle).Methods("GET")

	routes.HandleFunc("/api/subscription/{id:[0-9]+}", apiServer.GetSubscriptionHandle).Methods("GET")
	routes.HandleFunc("/api/subscription/all", apiServer.GetAllSubscriptionsHandle).Methods("GET")

	routes.HandleFunc("/api/user/{id:[0-9]+}", apiServer.GetUserHandle).Methods("GET")
	routes.HandleFunc("/api/user/all", apiServer.GetAllUsersHandle).Methods("GET")

	//Rotas de criação
	routes.HandleFunc("/api/hackathonUser", apiServer.CreateUserHandle).Methods("POST")

	routes.HandleFunc("/api/hackathon", apiServer.CreateUserHandle).Methods("POST")

	routes.HandleFunc("/api/subscription", apiServer.CreateUserHandle).Methods("POST")

	routes.HandleFunc("/api/user", apiServer.CreateUserHandle).Methods("POST")

	http.Handle("/", routes)
	http.HandleFunc("/ws", chatHandler.HandleConnections)
	http.HandleFunc("/sendMessage", chatHandler.SendViaPost)

	go chatHandler.HandleMessages()

	log.Println("http server started on " + os.Getenv("PORT"))
	log.Println("database started on " + model.Connection())
	log.Println("bonsai started on " + model.ElasticSearch())

	error := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}

}
