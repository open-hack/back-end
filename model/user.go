package model

import (
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID           uint        `gorm:"primary_key" json:"id,omitempty"`
	ProfileImage []byte      `gorm:"column:profile_image" json:"profileImage,omitempty"`
	Username     string      `gorm:"column:username" json:"username,omitempty"`
	Name         string      `gorm:"column:name" json:"name,omitempty"`
	Password     string      `gorm:"column:password" json:"password,omitempty"`
	Email        string      `gorm:"column:email" json:"email,omitempty"`
	Minibio      string      `gorm:"column:minibio" json:"minibio,omitempty"`
	LinkedInURL  string      `gorm:"column:linkedin_url" json:"linkedinUrl,omitempty"`
	GithubURL    string      `gorm:"column:github_url" json:"githubUrl,omitempty"`
	FacebookURL  string      `gorm:"column:facebook_url" json:"facebookUrl,omitempty"`
	BehanceURL   string      `gorm:"column:behance_url" json:"behanceUrl,omitempty"`
	DribbbleURL  string      `gorm:"column:dribbble_url" json:"dribbbleUrl,omitempty"`
	InstagramURL string      `gorm:"column:instagram_url" json:"instagramUrl,omitempty"`
	TwitterURL   string      `gorm:"column:twitter_url" json:"twitterUrl,omitempty"`
	Profile      string      `gorm:"column:profile" json:"profile,omitempty"`
	CreatedOn    time.Time   `gorm:"column:created_on" json:"created_on,omitempty"`
	LastLogin    time.Time   `gorm:"column:last_login" json:"lastLogin,omitempty"`
	Hackathons   []Hackathon `gorm:"many2many:hackathon_user;"`
}

//CreateUser: criar um usuário
func (dsd *WeeHackDB) CreateUser(user *User) error {
	result := dsd.Db.Table("public.user").Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//GetUsers: retorna um usuário
func (dsd *WeeHackDB) GetUser(id int) (*User, error) {
	user := User{}

	result := dsd.Db.Table("public.user").Preload("Hackathons").First(&user, "id = ?", id)

	if result.Error != nil {
		log.Println("error on get data from user", result.Error)
		return nil, result.Error
	}
	return &user, nil
}

//GetUsers: retorna todos os usuários
func (dsd *WeeHackDB) GetAllUsers() (*[]User, error) {
	users := []User{}
	result := dsd.Db.Table("public.user").Preload("Hackathons").Find(&users)
	if result.Error != nil {
		log.Println("error on get data from user", result.Error)
		return nil, result.Error
	}
	return &users, nil
}
