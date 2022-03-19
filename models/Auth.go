package models

type AutoLogin struct {
	Code string `example:"xtdad-fdfdsf"`
}

type Token struct {
	AccessToken string `example:"cdaefj93sds-eqedsdsdsa-3sadasdsss"`
	ExpiresIn   int    `example:"186624000"`
	TokenType   string `example:"bearer"`
}

type UserProfileForm struct {
	Code          string `example:"xtdad-fdfdsf"`
	EncryptedData string `example:"7SfFtStsHqKZYhbIkke3BH2bCRzGD15T0jEiUtuksrl9lDeHm9LsPmswJymBXuinPCiXkZhd/uq7s7pACTvbWuvvoKEwz5fAJ6Vr9bTx79XVxiIN4r+Fwm6QHO9DjPkFrxTGAZvMYLyH6IOyOV/nmmlMoBM3G4peSnBi1qCYukwlyCMNp67lb93wSiPAoI7eRhYYw8ayPTsZ/MAJ9CBBUiCwM5aFOUWrMKNTikeq7YVjNCv7KCz0LJTrMKda0YMS0J/034L8x9vJ1OnIkxlWVMQEy/f55IfWVHI1I1fSKd5azzyVKXCbWDpU0PLJnU8XM/l4L7ZUlDOcRMR5KQVGhB9rIjVkykdXUPQK87v8lpnitslK06XceOJqDjK6mRkhJWOYpFUozZa6idFV6xmLZX8bkBsLxczzp1h/satEH7rIz3nKbxd3O1c+3dI2soSt8qFtaumcGdwhenTm+at0gxccAp8JD8PZiB5ZDLTofZIQ4RmI004SIExYUDZUje9mZO+3aC8McVwzrEyK7NKD/NZ5/dYPgDRwzBl1Vm99niY="`
	UserInfo      string `example:""`
	Iv            string `example:"z3tGYrgMcbLzd0qXqZuduQ=="`
	RawData       string `example:""`
	Signature     string `example:""`
}
