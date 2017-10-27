package main

// DishRepository contains methods to query and mutate Dishs in the database
type DishRepository struct {
}

// NewDishRepo gets a new instance of a Dish repo
func NewDishRepo() *DishRepository {
	return &DishRepository{}
}

// getAllDishes returns all available Dish items
func (r *DishRepository) getAllDishes() ([]*Dish, error) {
	dishes := []*Dish{}

	err := DB.Select(&dishes, `SELECT * from dish;`)
	return dishes, err
}

// getDish retrieves a Dish with the given ID
func (r *DishRepository) getDish(id int) (*Dish, error) {
	dish := Dish{}

	err := DB.Get(&dish, `SELECT * from dish WHERE ID = $1;`, id)
	return &dish, err
}

// createDish create a Dish object in the repository
func (r *DishRepository) createDish(dish *Dish) {
	DB.NamedExec(`INSERT INTO dish
		(
			name,
			location_id
		)
		values
		(
			:name,
			:location_id
		);
	`, dish)
}

// updateDish updates an existing Dish
func (r *DishRepository) updateDish(dish *Dish) {
	DB.NamedExec(`UPDATE dish SET
		name = :name,
		location_id = :location_id
	WHERE ID = :id;`, dish)
}
