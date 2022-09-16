package models

type FocusRequest struct {
	On     string `json:"on" example:"2"`
	Status string `json:"status" example:"1"`
}

type Fan struct {
	Id         int    `json:"uid" example:"1"`
	UserName   string `json:"username" example:"苹果香蕉"`
	PetName    string `json:"petName" example:"一个宠物"`
	Age        int    `json:"age" example:"26"`
	Location   string `json:"location" example:"广东，深圳"`
	Occupation string `json:"occupation" example:"产品经理"`
	Images     string `json:"coverImage" example:"https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTL13ic0iaA0ffWldrLjv9Ou02CuJCcjuKJ7rAzatVEzEUsrceUEdIuSiaR7bnicf5X2puMFRNDLrPEJlw/132"`
	Status     int    `json:"status" example:"1"`
}

type ViewRequest struct {
	On string `json:"on" example:"2"`
}

type View struct {
	Id         int    `json:"uid" example:"1"`
	UserName   string `json:"username" example:"香蕉苹果"`
	PetName    string `json:"petName" example:"一个宠物"`
	Age        int    `json:"age" example:"26"`
	Location   string `json:"location" example:"广东，深圳"`
	Occupation string `json:"occupation" example:"产品经理"`
	Images     string `json:"coverImage" example:"https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTL13ic0iaA0ffWldrLjv9Ou02CuJCcjuKJ7rAzatVEzEUsrceUEdIuSiaR7bnicf5X2puMFRNDLrPEJlw/132"`
	Status     int    `json:"status" example:"1"`
	Message    string `json:"message" example:"Ta好像对你很感兴趣"`
	Highlight  string `json:"highlight" example:"感兴趣"`
}

type ViewTo struct {
	Id         int    `json:"uid" example:"1"`
	UserName   string `json:"username" example:"香蕉苹果"`
	PetName    string `json:"petName" example:"一个宠物"`
	Age        int    `json:"age" example:"26"`
	Location   string `json:"location" example:"广东，深圳"`
	Occupation string `json:"occupation" example:"产品经理"`
	Images     string `json:"coverImage" example:"https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTL13ic0iaA0ffWldrLjv9Ou02CuJCcjuKJ7rAzatVEzEUsrceUEdIuSiaR7bnicf5X2puMFRNDLrPEJlw/132"`
	Status     int    `json:"status" example:"1"`
}
