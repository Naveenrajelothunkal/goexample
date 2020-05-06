package skillresolvers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"

	"graphql-go-crud/data/connector/mongoConnector"
	"graphql-go-crud/data/model/mongoModel/skilldetailsmodel"
)

func SkillListResolver(s graphql.ResolveParams) (interface{}, error) {
	client := mongoConnector.MongoConnector()
	collection := client.Database("skilltest").Collection("skillInfo")
	cursor, err := collection.Find(context.Background(), bson.M{})
	var skillData skilldetailsmodel.SkillType
	if err != nil {
		log.Fatal(err)
	}
	var skillList []bson.M
	if err = cursor.All(context.Background(), &skillList); err != nil {
		log.Fatal()
	}
	var skillSlice []skilldetailsmodel.SkillType

	for i := 0; i < len(skillList); i++ {
		list := skillList[i]
		data, _ := json.Marshal(list)
		json.Unmarshal(data, &skillData)
		skillSlice = append(skillSlice, skillData)
	}
	return skillSlice, nil
}
