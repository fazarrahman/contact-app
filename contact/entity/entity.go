package entity

import "time"

type Contacts struct {
	Id        string     `db:"id"`
	Name      string     `db:"name"`
	Gender    string     `db:"gender"`
	Phone     string     `db:"phone"`
	Email     string     `db:"email"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
