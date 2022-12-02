package domain

type User struct {
	ModelBase
	Email     string
	FirstName string
	LastName  string
}
