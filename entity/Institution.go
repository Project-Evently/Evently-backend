package entity

type Institute struct {
	InstituteID   int    `json:"InstituteID"`
	InstituteName string `json:"InstituteName"`
	AdminId       string `json:"AdminId"`
	Password      string `json:"Password"`
}
