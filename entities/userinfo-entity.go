package entities

import "time"

// UserInfo .
type UserInfo struct {
	UID        int64     `xorm:"pk autoincr"`
	UserName   string    `xorm:"notnull unique"`
	DepartName string    `xorm:"notnull"`
	Created    time.Time `xorm:"created"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	return &u
}
