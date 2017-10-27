package main

// ReviewRepository contains methods to query and mutate Reviews in the database
type ReviewRepository struct {
}

// NewReviewRepo gets a new instance of a Review repo
func NewReviewRepo() *ReviewRepository {
	return &ReviewRepository{}
}

// getAllReviews returns all available Review items
func (r *ReviewRepository) getAllReviews() ([]*Review, error) {
	reviews := []*Review{}

	err := DB.Select(&reviews, `SELECT * from review;`)
	return reviews, err
}

// getReview retrieves a Review with the given ID
func (r *ReviewRepository) getReview(id int) (*Review, error) {
	review := Review{}

	err := DB.Get(&review, `SELECT * from review WHERE ID = $1;`, id)
	return &review, err
}

// createReview create a Review object in the repository
func (r *ReviewRepository) createReview(review *Review) error {
	_, err := DB.NamedExec(`INSERT INTO review
		(
			user_id,
			dish_id,
			title,
			rating,
			body,
			date_created,
			last_edited
		)
		values
		(
			:user_id,
			:dish_id,
			:title,
			:rating,
			:body,
			NOW(),
			NOW()
		);
	`, review)
	return err
}

// updateReview updates an existing Review
func (r *ReviewRepository) updateReview(review *Review) error {
	_, err := DB.NamedExec(`UPDATE Review SET
		user_id = :user_id,
		dish_id = :dish_id,
		title = :title,
		rating = :rating,
		body = :body,
		last_edited = NOW()
	WHERE ID = :id;`, review)
	return err
}
