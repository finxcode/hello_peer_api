package models

type RelationStat struct {
	Know_Me_Total    int `json:"know_me_total" example:"20"`
	Know_Me_New      int `json:"know_me_new" example:"3"`
	Focus_On_Total   int `json:"focus_on_total" example:"24"`
	Focused_By_Total int `json:"focused_by_total" example:"50"`
	Focus_By_New     int `json:"focus_by_new" example:"4"`
	Viewed_By_Total  int `json:"viewed_by_total" example:"200"`
	Viewed_By_New    int `json:"viewed_by_new" example:"10"`
}
