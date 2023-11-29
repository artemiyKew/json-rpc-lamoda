package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *PostgresDB {
	return &PostgresDB{
		db: db,
	}
}

func NewDB(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return &sqlx.DB{}, err
	}

	if err := db.Ping(); err != nil {
		return &sqlx.DB{}, err
	}
	return db, nil
}

func (p *PostgresDB) Close() error {
	if p.db.Ping() != nil {
		return p.db.Close()
	}
	return nil
}
