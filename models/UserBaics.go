package models

import (
	"gorm.io/gorm"
)
type UserBasic struct{
    gorm.Model
	Name string
	PassWord string
	Phone string
	Email string
	Identity string
	ClientIp string
	ClientPort string
	LoginTime uint64
	HeartbeatTime uint64
	LogOutTime uint64
	IsLogOut bool
	DeviceInfo string
}
func (table *UserBasic) TableName() string{
	return "user_basic"
}