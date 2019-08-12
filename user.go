package gomonorepo

type User struct {
	ID   string
	Name string
	Age  int
}

// NewUser creates a new user
func NewUser(id, name string, age int) *User {
	return &User{
		ID:   id,
		Name: name,
		Age:  age,
	}
}
