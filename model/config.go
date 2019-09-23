package model

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// Modelo de acesso ao banco
type WeeHackDB struct {
	Db *gorm.DB
}

//Inicialização do repository
func (dsd *WeeHackDB) MustInit() {
	var err error
	dsd.Db, err = gorm.Open("postgres", connection())
	if err != nil {
		log.Println("error on connect database")
		return
	}
}

//Helper para formatar string de conexão
func connection() string {
	return os.Getenv("DATABASE_URL")
}

func elasticSearch() string {
	return os.Getenv("BONSAI_URL")
}
