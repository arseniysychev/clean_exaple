package models

type User struct {
	ModelBase
	Email     string
	FirstName string
	LastName  string
}
