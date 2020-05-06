package user_resolvers

import (
	"context"
	"encoding/json"
	"graphql-go-crud/data/connector/mongoConnector"
	"graphql-go-crud/data/model/userInfoModel/userModel"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/graphql-go/graphql"
)

func UserByIDResolver(p graphql.ResolveParams) (interface{}, error) {
	client := mongoConnector.MongoConnector()
	collection := client.Database("UserInfo").Collection("UserData")
	id, ok := p.Args["id"].(int)
	var userData userModel.UserDetails
	if ok {
		filterCursor, err := collection.Find(context.Background(), bson.M{"id": id})
		if err != nil {
			log.Fatal(err)
		}
		var userList []bson.M
		if err = filterCursor.All(context.Background(), &userList); err != nil {
			log.Fatal(err)
		}
		userMap := userList[0]
		data, _ := json.Marshal(userMap)
		userResult := json.Unmarshal(data, &userData)
		return userData, userResult
	}
	return nil, nil
}
