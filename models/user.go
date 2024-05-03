package models

import (
	"time"
	"gorm.io/gorm"
)
type User struct{
    gorm.Model
	Name string `json:"name"`
	PassWord string `json:"password"`
	Salt string `json:"salt"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Identity string `json:"identity"`
	ClientIp string `json:"clientip"`
	ClientPort string `json:"clientport"`
	LoginTime time.Time `json:"logintime"`
	HeartbeatTime time.Time `json:"heartbeattime"`
	LoginOutTime time.Time   `gorm:"column:login_out_time" json:"loginouttime"`
	IsLogOut bool `json:"islogout"`
	DeviceInfo string `json:"deviceinfo"`
}
func NewUserWithCurrentTime(name string,password string,salt string) *User{
     return &User{
		Name:name,
		PassWord:password,
		Salt: salt,
        LoginTime:time.Now(),
		HeartbeatTime: time.Now(),
		LoginOutTime: time.Now(),
	 }
}
func NewUser() *User{
	return &User{}
}
func NewUserWithoutTime(name string,password string) *User{
	return &User{
		Name:name,
		PassWord:password,
	 }
}
func (table *User) TableName() string{
	return "user_info"
}

