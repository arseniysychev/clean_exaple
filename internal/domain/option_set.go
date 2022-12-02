package domain

type OptionSet struct {
	ModelBase
	Project Project
	Title   string
	Public  bool
}

type Option struct {
	ModelUUID
	Position uint8
}
