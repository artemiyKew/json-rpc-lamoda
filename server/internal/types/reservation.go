package types

type Reservation struct {
	ID          int    `db:"id"`
	UniqueCode  string `db:"unique_code"`
	WarehouseID int    `db:"warehouse_id"`
	Quantity    int    `db:"quantity"`
	Status      string `db:"status"`
}

const (
	ReservationStatusReserved = "reserved"
	ReservationStatusCanceled = "canceled"
)
