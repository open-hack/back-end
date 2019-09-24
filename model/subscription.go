package model

import "log"

type Subscription struct {
	ID                 int64  `gorm:"column:id;primary_key:true" json:"id,omitempty"`
	UserID             int64  `gorm:"column:user_id" json:"userId,omitempty"`
	CPF                string `gorm:"column:cpf" json:"cpf,omitempty"`
	RG                 string `gorm:"column:rg" json:"rg,omitempty"`
	Telephone          string `gorm:"column:telephone" json:"telephone,omitempty"`
	State              string `gorm:"column:state" json:"state,omitempty"`
	City               string `gorm:"column:city" json:"city,omitempty"`
	Address            string `gorm:"column:address" json:"address,omitempty"`
	Gender             string `gorm:"column:gender" json:"gender,omitempty"`
	Race               string `gorm:"column:race" json:"race,omitempty"`
	University         string `gorm:"column:university" json:"university,omitempty"`
	Course             string `gorm:"column:course" json:"course,omitempty"`
	GraduationYear     int    `gorm:"column:graduation_year" json:"graduationYear,omitempty"`
	EnglishLevel       string `gorm:"column:english_level" json:"englishLevel,omitempty"`
	Deficiency         string `gorm:"column:deficiency" json:"deficiency,omitempty"`
	FoodRestriction    string `gorm:"column:food_restriction" json:"foodRestriction,omitempty"`
	Allergy            string `gorm:"column:allergy" json:"allergy,omitempty"`
	SomethingImportant string `gorm:"column:something_important" json:"somethingImportant,omitempty"`
	ProudProject       string `gorm:"column:proud_project" json:"proudProject,omitempty"`
	ProjectLinks       string `gorm:"column:project_links" json:"projectLinks,omitempty"`
	HasBeenToHackathon string `gorm:"column:has_been_to_hackathon" json:"hasBeenToHackathon,omitempty"`
	HasBeenDescription string `gorm:"column:has_been_description" json:"hasBeenDescription,omitempty"`
	Profile            string `gorm:"column:profile" json:"profile,omitempty"`
	HowDidDiscover     string `gorm:"column:how_did_discover" json:"howDidDiscover,omitempty"`
	ShirtSize          string `gorm:"column:shirt_size" json:"shirtSize,omitempty"`
	Agreement          string `gorm:"column:agreement" json:"agreement,omitempty"`
	CompanyName        string `gorm:"column:company_name" json:"companyName,omitempty"`
	IsWorking          string `gorm:"column:is_working" json:"isWorking,omitempty"`
	HabilityOneLevel   int    `gorm:"column:hability_one_level" json:"habilityOneLevel,omitempty"`
	HabilityTwoLevel   int    `gorm:"column:hability_two_level" json:"habilityTwoLevel,omitempty"`
	HabilityThreeLevel int    `gorm:"column:hability_three_level" json:"habilityThreeLevel,omitempty"`
	HabilityFourLevel  int    `gorm:"column:hability_four_level" json:"habilityFourLevel,omitempty"`
}

//CreateSubscription: criar uma inscrição
func (dsd *WeeHackDB) CreateSubscription(subscription *Subscription) error {
	result := dsd.Db.Table("public.subscriptions").Create(subscription)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//GetSubscription: retorna uma inscrição
func (dsd *WeeHackDB) GetSubscription(id int) (*Subscription, error) {
	subscription := Subscription{}
	result := dsd.Db.Table("public.subscriptions").Where("id = ?", id).First(&subscription)
	if result.Error != nil {
		log.Println("error on get data from subscription", result.Error)
		return nil, result.Error
	}
	return &subscription, nil
}

//GetUsers: retorna todas as inscrições
func (dsd *WeeHackDB) GetAllSubscriptions() (*[]Subscription, error) {
	subscriptions := []Subscription{}
	result := dsd.Db.Table("public.subscriptions").Find(&subscriptions)
	if result.Error != nil {
		log.Println("error on get data from subscription", result.Error)
		return nil, result.Error
	}
	return &subscriptions, nil
}
