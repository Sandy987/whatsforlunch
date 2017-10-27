package main

// LocationRepository contains methods to query and mutate Locations in the database
type LocationRepository struct {
}

// NewLocationRepo gets a new instance of a Location repo
func NewLocationRepo() *LocationRepository {
	return &LocationRepository{}
}

// getAllLocations returns all available Location items
func (r *LocationRepository) getAllLocations() ([]*Location, error) {
	Locations := []*Location{}

	err := DB.Select(&Locations, `SELECT * from location;`)
	return Locations, err
}

// getLocation retrieves a Location with the given ID
func (r *LocationRepository) getLocation(id int) (*Location, error) {
	Location := Location{}

	err := DB.Get(&Location, `SELECT * from location WHERE ID = $1;`, id)
	return &Location, err
}

// createLocation create a Location object in the repository
func (r *LocationRepository) createLocation(loc *Location) error {
	_, err := DB.NamedExec(`INSERT INTO location
		(
			name
		)
		values
		(
			:name
		);
	`, loc)
	return err
}

// updateLocation updates an existing Location
func (r *LocationRepository) updateLocation(loc *Location) error {
	_, err := DB.NamedExec(`UPDATE location SET
		name = :name
	WHERE ID = :id;`, loc)
	return err
}
