package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // users is a slice that holds a collection pointers to Users
	nextID = 1     // simulate a primary key //* NOTE: at package level we don't need impicit initizlization with ':='
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
