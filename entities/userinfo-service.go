package entities

type UserInfoAtomicService struct{}

var UserInfoService = UserInfoAtomicService{}

func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := engine.Insert(u)
	checkErr(err)
	return err
}

func (*UserInfoAtomicService) FindAll() []UserInfo {
	everyone := make([]UserInfo, 0)
	err := engine.Find(&everyone)
	checkErr(err)
	return everyone
}

func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	id64 := int64(id)
	userQuery := &UserInfo{UID: id64}
	_, err := engine.Get(userQuery)
	checkErr(err)
	return userQuery
}
