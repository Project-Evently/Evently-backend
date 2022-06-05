package entity

type User struct {
	UserId          int    `json:"UserId"`
	Username        string `json:"Username"`
	Password        string `json:"Password"`
	CollegeName     string `json:"CollegeName"`
	BranchName      string `json:"BranchName"`
	CurrentYear     string `json:"CurrentYear"`
	ContactNumber   string `json:"ContactNumber"`
	Email           string `json:"Email"`
	GithubId        string `json:"GithubId"`
	UniqueStudentId string `json:"UniqueStudentId"`
	EventRegistered []int  `json:"EventRegistered"`
	//Social          []string `json:"SocialLinks"`
}

type UpdatePassword struct {
	UniqueStudentId string `json:"UniqueStudentId"`
	Password        string `json:"Password"`
}
