package dynamodb

import (
	"fmt"
	"log/slog"
)

type (
	baseDynamoDbProvider struct {
		tableName string
	}

	IBaseDynamoDbProvider interface {
		GetAllItemsByPk()
		InsertItems()
	}
)

func (provider *baseDynamoDbProvider) GetAllItemsByPk() {
	slog.Info(fmt.Sprintf("dummy GetAllItemsByPk working on table ---> %s", provider.tableName))
}

func (provider *baseDynamoDbProvider) InsertItems() {
	slog.Info(fmt.Sprintf("dummy InsertItems working on table ---> %s", provider.tableName))
}
