package types

type Shipping struct {
	ID          int    `db:"id"`
	UniqueCode  string `db:"unique_code"`
	WarehouseID int    `db:"warehouse_id"`
	Quantity    int    `db:"quantity"`
}
