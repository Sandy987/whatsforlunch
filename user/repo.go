package user

import "github.com/Sandy987/whatsforlunch/domain"

// GetAllUsers returns all available User items
func GetAllUsers() (Users, error) {
	users := Users{}

	err := domain.DB.Select(&users, "SELECT * from \"user\"")
	return users, err
}

// GetUser retrieves a User with the given ID
func GetUser(id int) (*User, error) {
	user := User{}

	err := domain.DB.Get(&user, "SELECT * from \"user\" WHERE ID = $1", id)
	return &user, err
}

// CreateUser create a User object in the repository
func CreateUser(u *HashedUser) {
	domain.DB.NamedExec(`INSERT INTO \"user\" 
		(
			username,
			first_name,
			last_name,
			email,
			password_hash,
		)
		values
		(
			:username,
			:first_name,
			:last_name,
			:email,
			:password_hash,
		)
	`, u)
}

// UpdateUser updates an existing User
func UpdateUser(u *User) {
	domain.DB.NamedExec(`UPDATE \"user\" SET
		first_name = :first_name,
		last_name = :last_name,
	WHERE ID = :id`, u)
}
