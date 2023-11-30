package postgres

import (
	"context"

	"github.com/artemiyKew/json-rpc-lamoda/internal/types"
	"github.com/sirupsen/logrus"
)

type ReservationRepo struct {
	*PostgresDB
}

func NewReservationRepo(db *PostgresDB) *ReservationRepo {
	return &ReservationRepo{db}
}

func (r *ReservationRepo) CreateReservation(ctx context.Context, reservation types.Reservation) error {
	query := `INSERT INTO Reservation (unique_code, warehouse_id, quantity, status) VALUES ($1, $2, $3, $4)`
	if err := r.db.QueryRowContext(ctx, query, reservation.UniqueCode, reservation.WarehouseID, reservation.Quantity, reservation.Status).Err(); err != nil {
		logrus.Fatalf("ReservationRepo:CreateReservation %s", err)
		return err
	}
	return nil
}

func (r *ReservationRepo) CancelReservation(ctx context.Context, reservation types.Reservation) ([]types.Reservation, error) {
	query := `UPDATE Reservation SET status = $1 WHERE unique_code = $2 AND status = $3 RETURNING *`
	reservations := make([]types.Reservation, 0)
	if err := r.db.SelectContext(ctx, &reservations, query, types.ReservationStatusCanceled, reservation.UniqueCode, types.ReservationStatusReserved); err != nil {
		logrus.Fatalf("ReservationRepo:CancelReservation %s", err)
		return nil, nil
	}

	return reservations, nil
}

func (r *ReservationRepo) GetAllReservations(ctx context.Context) ([]types.Reservation, error) {
	query := `SELECT * FROM Reservation`
	reservations := make([]types.Reservation, 0)

	if err := r.db.SelectContext(ctx, &reservations, query); err != nil {
		logrus.Fatalf("ReservationRepo:GetAllReservations %s", err)
		return nil, err
	}

	return reservations, nil
}
