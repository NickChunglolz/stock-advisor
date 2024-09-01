package usecase

import (
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/client"
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/infrastructure/repository"
)

type StockCommand struct {
	client *client.TwseApiClient
	repo   *repository.StockRepository
}

func NewStockCommand(client *client.TwseApiClient, repo *repository.StockRepository) *StockCommand {
	return &StockCommand{
		client: client,
		repo:   repo,
	}
}

func (command *StockCommand) GetAndStoreStocks() error {
	replies, err := command.client.GetStockInfos(mockQueryString)
	if err != nil {
		return err
	}

	return command.repo.Save(replies)
}
