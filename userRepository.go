package main

// UserRepository contains methods to query and mutate users in the database
type UserRepository struct {
}

// UserRepo is a global instance of this struct
// TODO: Is there a better way of doing this?
var UserRepo = UserRepository{}

// GetAllUsers returns all available User items
func (r *UserRepository) GetAllUsers() ([]*User, error) {
	users := []*User{}

	err := DB.Select(&users, "SELECT * from \"user\"")
	return users, err
}

// GetUser retrieves a User with the given ID
func (r *UserRepository) GetUser(id int) (*User, error) {
	user := User{}

	err := DB.Get(&user, "SELECT * from \"user\" WHERE ID = $1", id)
	return &user, err
}

// CreateUser create a User object in the repository
func (r *UserRepository) CreateUser(u *HashedUser) {
	DB.NamedExec(`INSERT INTO \"user\" 
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
func (r *UserRepository) UpdateUser(u *User) {
	DB.NamedExec(`UPDATE \"user\" SET
		first_name = :first_name,
		last_name = :last_name,
	WHERE ID = :id`, u)
}
