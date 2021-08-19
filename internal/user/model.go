package user

type User struct {
	Resource user
}

func NewUser(resource user) *User {
	return &User{
		Resource: resource,
	}
}
