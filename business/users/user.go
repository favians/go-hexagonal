package users

type User struct {
	Id          string
	Email       string
	Password    string
	Name        string
	CurrentRoom string
}

func NewUser(
	id string,
	email string,
	password string,
	name string,
	currentRoom string) User {

	return User{
		Id:          id,
		Email:       email,
		Password:    password,
		Name:        name,
		CurrentRoom: currentRoom,
	}
}