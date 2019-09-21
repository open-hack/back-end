package model

import(
	"time"
)

type Config struct {
    ElasticSearchUrl string `env:"ELASTICSEARCH_URL" envDefault:"Slomek"`
}

type Message struct {
	ID			string			`json:"_id,omitempty"`
	Text		string			`json:"text"`
	CreatedAt	time.Time		`json:"createdAt"`
	User		User			`json:"user"`
}

type User struct {
	ID			int				`json:"_id,omitempty"`
	Name		string			`json:"name"`
	Avatar		string			`json:"avatar"`
}