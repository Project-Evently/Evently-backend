package entity

type Event struct {
	EventId        int    `json:"eventId"`
	ClubId         int    `json:"clubId"`
	InstituteId    string `json:"instituteId"`
	Description    string `json:"description"`
	EventDateIST   string `json:"eventDateIST"`
	EventTimeIST   string `json:"eventTimeIST"`
	EventLocation  string `json:"eventLocation"`
	EventOrganizer string `json:"eventOrganizer"`
	EventContact   string `json:"eventContact"`
	EventLink      string `json:"eventLink"`
}
