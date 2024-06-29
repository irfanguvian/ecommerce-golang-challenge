package models

type AccessToken struct {
	UserID    uint
	AccessID  string
	ExpiredAt int64
}
