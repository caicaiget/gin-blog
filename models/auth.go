package models

type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool, int) {
	var auth Auth
	db.Select("id").Where(Auth{Username : username, Password : password}).First(&auth)
	if auth.ID > 0 {
		return true, auth.ID
	}

	return false, auth.ID
}

func GetUserById(id int) (auth Auth) {
	db.Where("id = ?", id).First(&auth)
	return
}
