package mongoResolvers

import (
	"graphql-go-crud/data/model/mongoModel/toolDetailsModel"
	"graphql-go-crud/data/services/mongoDbServices"

	"github.com/graphql-go/graphql"
)

func InsertOneMongoDb(params graphql.ResolveParams) (interface{}, error) {
	data := toolDetailsModel.ToolDetails{
		Order_Id:   params.Args["Order_Id"].(string),
		Created_On: params.Args["Created_On"].(string),
		Updated_On: params.Args["Updated_On"].(string),
	}
	mongoDbServices.InsertOne(data)
	return data, nil
}
