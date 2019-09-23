package main

import (
	"log"
	"net/http"
	"os"

	"github.com/caarlos0/env"
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

	var cfg model.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Failed to parse ENV")
	}

	elasticClient, err := elastic.NewClient(elastic.SetURL(cfg.ElasticSearchUrl), elastic.SetSniff(false))
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

	fs := http.FileServer(http.Dir("../public"))

	//Rotas de consulta
	routes.HandleFunc("/api/player/{id:[0-9]+}", apiServer.GetUserHandle).Methods("GET")
	routes.HandleFunc("/api/player/all", apiServer.GetAllUsersHandle).Methods("GET")

	//Rotas de criação
	routes.HandleFunc("/api/player", apiServer.CreateUserHandle).Methods("POST")

	http.Handle("/", fs)
	http.HandleFunc("/ws", chatHandler.HandleConnections)
	http.HandleFunc("/sendMessage", chatHandler.SendViaPost)

	go chatHandler.HandleMessages()

	log.Println("http server started on " + os.Getenv("PORT"))

	error := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}

}
