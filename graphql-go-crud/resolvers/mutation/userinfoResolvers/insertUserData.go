package userinfoResolvers

import (
	"graphql-go-crud/data/services/userInfoServices"

	"github.com/graphql-go/graphql"

	"graphql-go-crud/data/model/userInfoModel/userModel"
)

func UserinfoMongoDb(params graphql.ResolveParams) (interface{}, error) {
	data := userModel.UserDetails{
		ID:         params.Args["id"].(int),
		Name:       params.Args["Name"].(string),
		Place:      params.Args["Place"].(string),
		Created_On: params.Args["Created_On"].(string),
		Updated_On: params.Args["Updated_On"].(string),
	}
	userInfoServices.InsertUserData(data)
	return data, nil
}
