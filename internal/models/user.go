package models

type User struct {
	Name      string
	Age       uint16
	Money     int16
	AvgGrades float64
	Happiness float64
	Hobbies   []string
}

func NewUser() *User {
	return &User{
		Hobbies: make([]string, 0),
	}
}
