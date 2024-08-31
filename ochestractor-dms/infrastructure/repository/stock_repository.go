package repository

import (
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/client"
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/utils"
	"github.com/go-pg/pg/v10"
)

type StockRepository struct {
	db *pg.DB
}

func NewStockRepository() *StockRepository {
	return &StockRepository{
		db: utils.NewDatabaseFactory().GetPostgres(),
	}
}

// TODO - Needs to create Stock's data model.
func (repo *StockRepository) Save(data []client.StockInfoReply) error {
	_, err := repo.db.Model(data).Insert()
	return err
}