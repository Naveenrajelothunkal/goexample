package mongoDbServices

import (
	"context"
	"log"
	
	"graphql-go-crud/data/connector/mongoConnector"
	"graphql-go-crud/data/model/mongoModel/toolDetailsModel"
)

func InsertOne(d toolDetailsModel.ToolDetails) bool {
	client := mongoConnector.MongoConnector()
	collection := client.Database("cloveDb").Collection("toolDetails")
	_, err := collection.InsertOne(context.Background(), d)
	if err != nil { 
		log.Fatal(err)
		return false
	}
	return true
}