package mongoDbServices

import (
	"context"
	"graphql-go-crud/data/connector/mongoConnector"
	"graphql-go-crud/data/model/mongoModel/skilldetailsmodel"
	"log"
)

func InsertSkillToMongo(s skilldetailsmodel.SkillType) bool {
	client := mongoConnector.MongoConnector()
	collection := client.Database("skilltest").Collection("skillInfo")
	_, err := collection.InsertOne(context.Background(), s)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
