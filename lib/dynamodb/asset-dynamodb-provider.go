package dynamodb

type assetDynamoDbProvider struct {
	baseDynamoDbProvider
}

func NewAssetDynamoDbProvider() IBaseDynamoDbProvider {
	return &assetDynamoDbProvider{
		baseDynamoDbProvider{
			tableName: "dev-asset-table",
		},
	}
}
