package model

import "log"

type Hackathon struct {
	ID         int64  `gorm:"primary_key" json:"id,omitempty"`
	Image      []byte `gorm:"column:image" json:"image,omitempty"`
	Title      string `gorm:"column:title" json:"title,omitempty"`
	Onboarding string `gorm:"column:onboarding" json:"onboarding,omitempty"`
	Users      []User `gorm:"many2many:hackathon_user;association_autocreate:false" json:"users,omitempty"`
}

//CreateHackathon: criar um hackathon
func (dsd *WeeHackDB) CreateHackathon(hackathon *Hackathon) error {
	result := dsd.Db.Table("public.hackathon").Create(hackathon)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//GetScouts: retorna um hackathon
func (dsd *WeeHackDB) GetHackathon(id int) (*Hackathon, error) {
	hackathon := Hackathon{}
	users := []User{}

	result := dsd.Db.Table("public.hackathon").First(&hackathon, "id = ?", id)

	if result.Error != nil {
		log.Println("error on get data from hackathon", result.Error)
		return nil, result.Error
	} else {

		dsd.Db.Model(&hackathon).Related(&users, "User")
		log.Println(users)
		hackathon.Users = users
	}
	return &hackathon, nil
}

//GetUsers: retorna todos os hackathon
func (dsd *WeeHackDB) GetAllHackathons() (*[]Hackathon, error) {
	hackathons := []Hackathon{}

	result := dsd.Db.Table("public.hackathon").Preload("User").Find(&hackathons)

	log.Println(result)

	if result.Error != nil {
		log.Println("error on get data from hackathon", result.Error)
		return nil, result.Error
	}
	return &hackathons, nil
}
