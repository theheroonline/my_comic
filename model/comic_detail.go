package model

type ComicDetail struct {
	Id      int    `json:"id" gorm:"column:id;"`
	ComicId int    `json:"comic_id" gorm:"column:comic_id;"`
	Url     string `json:"url" gorm:"column:url;"`
}

func (ComicDetail) TableName() string {
	return "comic_detail"
}
