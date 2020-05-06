package userInfoServices

import (
	"context"
	"log"

	"graphql-go-crud/data/connector/mongoConnector"
	"graphql-go-crud/data/model/userInfoModel/userModel"
)

func InsertUserData(u userModel.UserDetails) bool {
	client := mongoConnector.MongoConnector()
	collection := client.Database("UserInfo").Collection("UserData")
	_, err := collection.InsertOne(context.Background(), u)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
