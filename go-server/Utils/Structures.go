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
