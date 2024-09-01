package repository

import "github.com/go-pg/pg/v10"

type Repositories struct {
	Stock *StockRepository
}

type Repository interface {
	Get() (any, error)
	Find() (any, error)
	Save() error
	Update() error
	Delete() (any, error)
}

func InitRepositories(db *pg.DB) *Repositories {
	return &Repositories{
		Stock: NewStockRepository(db),
	}
}
