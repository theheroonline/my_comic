package model

type Comic struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Icon string `json:"icon" gorm:"column:icon;"`
	Nub  int    `json:"nub" gorm:"column:nub;"`
}

func (Comic) TableName() string {
	return "comic"
}
