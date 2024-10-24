package entity

type (
	Product struct {
		ID          int    `db:"id"`
		Name        string `db:"name"`
		Description string `db:"description"`
		Price       int    `db:"price"`
		CreatedBy   int    `db:"created_by"`
	}
)
