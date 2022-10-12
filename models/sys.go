package models

type Agreement struct {
	Name    string `json:"name" example:"terms"`
	Title   string `json:"title" example:"hello peer用户服务协议"`
	Content string `json:"content" example:"协议正文"`
}
