package model

import "log"

type HackathonUser struct {
	ID          int64  `gorm:"column:hackathon_user_id" json:"id,omitempty"`
	State       string `gorm:"column:hackathon_state" json:"hackathonState,omitempty"`
	HackathonID int64  `gorm:"column:hackathon_id" json:"hackathonId,omitempty"`
	USerID      int64  `gorm:"column:user_id" json:"userId,omitempty"`
}

//CreateHackathonUser: criar associação de hackathon e hackathoner
func (dsd *WeeHackDB) CreateHackathonUser(hackathonUser *HackathonUser) error {
	result := dsd.Db.Table("public.hackathon_user").Create(hackathonUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//GetScouts: retorna uma associação de hackathon e hackathoner
func (dsd *WeeHackDB) GetHackathonUser(id int) (*HackathonUser, error) {
	hackathonUser := HackathonUser{}
	result := dsd.Db.Table("public.hackathon_user").Where("id = ?", id).First(&hackathonUser)
	if result.Error != nil {
		log.Println("error on get data from hackathon_user", result.Error)
		return nil, result.Error
	}
	return &hackathonUser, nil
}

//GetUsers: retorna todos os hackathon
func (dsd *WeeHackDB) GetAllHackathonUsers() (*[]HackathonUser, error) {
	var hackathonUsers []HackathonUser
	result := dsd.Db.Table("public.hackathon").Find(&hackathonUsers)
	if result.Error != nil {
		log.Println("error on get data from hackathonUser", result.Error)
		return nil, result.Error
	}
	return &hackathonUsers, nil
}
