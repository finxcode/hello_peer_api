package models

type UserGender struct {
	Gender int `example:"1"`
}

type BasicInfo struct {
	Gender        int     `example:"1"`
	Birth_day     string  `example:"2010-10-09"`
	Constellation string  `example:"白羊座"`
	Height        float32 `example:"178"`
	Weight        float32 `example:"56.5"`
	Education     string  `example:"本科"`
	Occupation    string  `example:"设计师"`
	Location      string  `example:"深圳"`
	Hometown      string  `example:"重庆"`
	Marriage      string  `example:"未婚"`
}

type HomepageInfo struct {
	UserName string       `json:"user_name" example:"用户1"`
	Location string       `json:"location" example:"广东,深圳"`
	Stat     RelationStat `json:"stat"`
	PetFood  int          `json:"pet_food" example:"320"`
	PetName  string       `json:"pet_name" example:"宠物猫猫"`
	Avatar   string       `json:"avatar" example:"avatar1-1234.jpg"`
}
