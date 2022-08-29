package models

type FocusRequest struct {
	On     string `json:"on" example:"2"`
	Status string `json:"status" example:"1"`
}
