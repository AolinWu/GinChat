package models

import (
	"time"

	"gorm.io/gorm"
)
type User struct{
    gorm.Model
	Name string
	PassWord string
	Phone string
	Email string
	Identity string
	ClientIp string
	ClientPort string
	LoginTime time.Time
	HeartbeatTime time.Time
	LoginOutTime time.Time   `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogOut bool
	DeviceInfo string
}
func (table *User) TableName() string{
	return "user"
}