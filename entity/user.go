package entity

type User struct {
	UserId          int      `json:"UserId"`
	Username        string   `json:"Username"`
	Password        string   `json:"Password"`
	UniqueStudentId string   `json:"UniqueStudentId"`
	EventRegistered []int    `json:"EventRegistered"`
	Social          []string `json:"SocialLinks"`
}
