package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Config struct {
	ElasticSearchUrl string `env:"BONSAI_URL"`
}

//Configuração de banco de dados, host,senha etc..
// Dados de produção são carregados via variaveis de ambiente (:
var Database struct {
	URL string `env:"DATABASE_URL"`
}

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
	return fmt.Sprintf(
		"%s?sslmode=require",
		Database.URL,
	)
}
