package types

type Warehouse struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	Availability bool   `db:"availability"`
}
