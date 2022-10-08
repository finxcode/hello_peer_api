package models

type UserSetting struct {
	Id           int    `json:"id" example:"1"`
	UserVerified int    `json:"verified" example:"1"`
	PetVerified  int    `json:"petVerified" example:"0"`
	HelloId      string `json:"helloId" example:"hp135895332"`
	Phone        string `json:"phone" example:"13726352437"`
	WechatName   string `json:"wechatName" example:"香蕉苹果"`
}
