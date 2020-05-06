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

func UserListResolver(p graphql.ResolveParams) (interface{}, error) {
	client := mongoConnector.MongoConnector()
	collection := client.Database("UserInfo").Collection("UserData")
	var userData userModel.UserDetails
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var userList []bson.M
	if err = cursor.All(context.Background(), &userList); err != nil {
		log.Fatal(err)
	}
	var userSlice []userModel.UserDetails
	for i := 0; i < len(userList); i++ {
		listdata := userList[i]
		data, _ := json.Marshal(listdata)
		json.Unmarshal(data, &userData)
		userSlice = append(userSlice, userData)
	}
	return userSlice, nil
}
