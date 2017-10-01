package user

// User contains information about a user
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email"`
}

// Users contains a collection of User items
type Users []User

// HashedUser should only be used internally as it contains the hash
type HashedUser struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	FirstName    string `json:"firstName" db:"first_name"`
	LastName     string `json:"lastName" db:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash" db:"password_hash"`
}
