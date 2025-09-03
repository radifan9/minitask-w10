package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type StoreRepository struct {
	db *pgxpool.Pool
}

func NewStoreRepository(db *pgxpool.Pool) *StoreRepository {
	return &StoreRepository{
		db: db,
	}
}
