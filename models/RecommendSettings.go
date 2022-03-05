package models

type RecommendSetting struct {
	Gender   string `example:"女生"`
	AgeMin   int    `example:"22"`
	AgeMax   int    `example:"40"`
	Location string `example:"只要同城"`
	Hometown string `example:"同省优先"`
	PetLover string `example:"喜欢就行"`
	Tags     string `example:"不限"`
}
