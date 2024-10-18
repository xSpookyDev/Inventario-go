package entity

type (
	User struct {
		ID       int    `db:"id" json:"id"`
		Email    string `db:"email" json:"email"`
		Name     string `db:"name" json:"name"`
		Password string `db:"password" json:"-"`
	}
)
