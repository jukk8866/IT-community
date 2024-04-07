package models

type User struct {
	UserId   int64
	Username string
	Password string
	AToken   string
	RToken   string
}

func (User) TableName() string {
	return "user"
}
