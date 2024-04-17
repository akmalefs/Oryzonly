package user 

type UserFormatter struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func UserFormat(user User, token string) UserFormatter{
	formatter := UserFormatter{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Token: token,
	}

	return formatter
}