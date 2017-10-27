package main

import "time"

// Temporarily keeping all the schema in a single file for cleanliness.
// If the schema gets really big, split it out.

// Location is a place.
type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Dish is a delectable (hopefully) meal.
type Dish struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LocationID int    `json:"locationId" db:"location_id"`
}

// Review is the meat and potatoes of this app.
type Review struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userId" db:"user_id"`
	DishID      int       `json:"dishId" db:"dish_id"`
	Title       string    `json:"title"`
	Rating      int       `json:"rating"`
	Body        string    `json:"body"`
	DateCreated time.Time `json:"dateCreated" db:"date_created"`
	LastEdited  time.Time `json:"lastEdited" db:"last_edited"`
}

// User contains information about a user
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email"`
}

// HashedUser should only be used internally as it contains the hash
type HashedUser struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	FirstName    string `json:"firstName" db:"first_name"`
	LastName     string `json:"lastName" db:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash" db:"password_hash"`
}
