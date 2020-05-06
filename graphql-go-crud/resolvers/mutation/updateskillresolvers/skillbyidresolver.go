package updateskillresolvers

import (
	"context"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"

	"graphql-go-crud/data/connector/mongoConnector"
)

func UpdateSkillResolver(s graphql.ResolveParams) (interface{}, error) {
	client := mongoConnector.MongoConnector()
	collection := client.Database("skilltest").Collection("skillInfo")
	skillId, ok := s.Args["skill_id"].(int)
	name, _ := s.Args["name"].(string)
	rating, _ := s.Args["rating"].(int)

	if ok {
		result, err := collection.ReplaceOne(context.Background(), bson.M{"skill_id": skillId}, bson.M{
			"skill_id": skillId,
			"name":     name,
			"rating":   rating,
		},
		)
		// result, err := collection.UpdateMany(
		// 	context.Background(),
		// 	bson.M{"skill_id": skillId},
		// 	bson.D{
		// 		{"$set", bson.D{{"name", name}, {"rating", rating}}},
		// 	},
		// )
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
		return result.ModifiedCount, nil
	}
	return nil, nil
}
