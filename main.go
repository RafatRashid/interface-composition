package main

import "interface-composition/lib/dynamodb"

func main() {
	assetDynamodbProvider := dynamodb.NewAssetDynamoDbProvider()
	gamePlayDynamodbProvider := dynamodb.NewGamePlayDynamoDbProvider()

	assetDynamodbProvider.GetAllItemsByPk()
	gamePlayDynamodbProvider.GetAllItemsByPk()
}
