package model

import "log"

type Hackathon struct {
	ID             int64   `gorm:"primary_key" json:"id,omitempty"`
	Image          []byte  `gorm:"column:image" json:"image,omitempty"`
	Title          string  `gorm:"column:title" json:"title,omitempty"`
	Onboarding     string  `gorm:"column:onboarding" json:"onboarding,omitempty"`
	HackathonState string  `gorm:"column:hackathon_state" json:"hackathonState,omitempty"`
	Users          []*User `gorm:"many2many:hackathon_user;" json:"users,omitempty"`
}

//CreateHackathon: criar um hackathon
func (dsd *WeeHackDB) CreateHackathon(hackathon *Hackathon) error {
	result := dsd.Db.Table("public.hackathons").Create(hackathon)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//GetScouts: retorna um hackathon
func (dsd *WeeHackDB) GetHackathon(id int) (*Hackathon, error) {
	hackathon := Hackathon{}

	result := dsd.Db.Table("public.hackathons").Preload("Users").First(&hackathon, "id = ?", id)

	if result.Error != nil && !result.RecordNotFound() {
		log.Println("error on get data from hackathon", result.Error)
		return nil, result.Error
	}
	return &hackathon, nil
}

//GetUsers: retorna todos os hackathon
func (dsd *WeeHackDB) GetAllHackathons() (*[]Hackathon, error) {
	hackathons := []Hackathon{}

	result := dsd.Db.Table("public.hackathons").Preload("Users").Find(&hackathons)

	log.Println(result)

	if result.Error != nil {
		log.Println("error on get data from hackathon", result.Error)
		return nil, result.Error
	}
	return &hackathons, nil
}
