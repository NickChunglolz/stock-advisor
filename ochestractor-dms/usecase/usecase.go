package usecase

import (
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/client"
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/repository"
)

type Usecases struct {
	StockQuery   *StockQuery
	StockCommand *StockCommand
}

func InitUsecases(cliets *client.Clients, repos *repository.Repositories) *Usecases {
	return &Usecases{
		StockQuery:   NewStockQuery(cliets.TwseApi),
		StockCommand: NewStockCommand(cliets.TwseApi, repos.Stock),
	}
}
