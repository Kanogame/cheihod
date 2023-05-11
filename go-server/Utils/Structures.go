package utils

type LoginJson struct {
	Username string
	Password string
}

type RegJson struct {
	Username string
	Password string
	Email    string
}

type NewToken struct {
	Token  string
	Userid int
}

type AddPlace struct {
	Name     string
	Place    string
	Date     string
	Time     string
	Capacity string
}

type TicketAdd struct {
	Token   string
	PlaceId string
}

type Ticket struct {
	UserId  string
	PlaceId string
}
