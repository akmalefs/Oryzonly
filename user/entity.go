package user

import "time"

type User struct {
	ID             int
	Name           string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	Occupation     string
	Phone          string
	Nationality    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
