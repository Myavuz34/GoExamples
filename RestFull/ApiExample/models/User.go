package models

//User struct
type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Profile   string
	Interests []Interest
}
