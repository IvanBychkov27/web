package models

type Contact struct {
	Name    string
	Address string
	EMail   string
}

func NewContact() *Contact {
	return &Contact{}
}
