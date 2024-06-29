package models

type AccessToken struct {
	UserID    uint `gorm:"primaryKey"`
	AccessID  string
	ExpiredAt int64
}
