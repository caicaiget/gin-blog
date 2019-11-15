package models

type Auth struct {
	ID int64 `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool, int64) {
	var auth Auth
	db.Select("id").Where(Auth{Username : username, Password : password}).First(&auth)
	if auth.ID > 0 {
		return true, auth.ID
	}

	return false, auth.ID
}

func GetUserById(id int64) (user Auth) {
	db.Where("id = ?", id).First(&user)
	return
}

func GetUserByName(name string) (user Auth, ok bool) {
	db.Where("name = ?", name).First(&user)
	if user.ID > 0 {
		return user, true
	}
	return user, false
}
