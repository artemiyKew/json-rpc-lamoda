package types

type Product struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Size       string `db:"size"`
	UniqueCode string `db:"unique_code"`
	Quantity   int    `db:"quantity"`
}
