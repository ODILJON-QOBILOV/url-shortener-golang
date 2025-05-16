package models

type Url struct {
	Id         int
	Url        string `gorm:"unique"`
	ShortedUrl string
}

type GetUrl struct {
	Url string `gorm:"unique"`
}