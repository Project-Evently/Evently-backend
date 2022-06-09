package entity

type Club struct {
	Id            int    `json:"id"`
	InstituteId   int    `json:"instituteId"`
	ClubName      string `json:"clubName"`
	ClubPresident string `json:"clubPresident"`
	AdminId       string `json:"adminId"`
	AdminPassword string `json:"adminPassword"`
}
