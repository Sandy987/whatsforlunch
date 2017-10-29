package main

// UserRepository contains methods to query and mutate users in the database
type UserRepository struct {
}

// NewUserRepo gets a new instance of a user repo
func NewUserRepo() *UserRepository {
	return &UserRepository{}
}

// getAllUsers returns all available User items
func (r *UserRepository) getAllUsers() ([]*User, error) {
	users := []*User{}

	err := DB.Select(&users, `SELECT * from "user";`)
	return users, err
}

// getUser retrieves a User with the given ID
func (r *UserRepository) getUser(id int) (*User, error) {
	user := User{}

	err := DB.Get(&user, `SELECT
		id,
		username,
		first_name,
		last_name,
		email
		from "user" WHERE ID = $1;`, id)
	return &user, err
}

func (r *UserRepository) getHashForUsername(username string) (string, error) {
	var hash string
	err := DB.Get(&hash, `SELECT
		password_hash
		from "user" WHERE username = $1;`, username)
	return hash, err
}

// createUser create a User object in the repository
func (r *UserRepository) createUser(u *HashedUser) error {
	_, err := DB.NamedExec(`INSERT INTO "user" 
		(
			username,
			first_name,
			last_name,
			email,
			password_hash
		)
		values
		(
			:username,
			:first_name,
			:last_name,
			:email,
			:password_hash
		);
	`, u)
	return err
}

// updateUser updates an existing User
func (r *UserRepository) updateUser(u *User) error {
	_, err := DB.NamedExec(`UPDATE "user" SET
		first_name = :first_name,
		last_name = :last_name
	WHERE ID = :id;`, u)
	return err
}
