package dynamodb

type gamePlayDynamoDbProvider struct {
	baseDynamoDbProvider
}

func NewGamePlayDynamoDbProvider() IBaseDynamoDbProvider {
	return &gamePlayDynamoDbProvider{
		baseDynamoDbProvider{
			tableName: "dev-game-play-table",
		},
	}
}
