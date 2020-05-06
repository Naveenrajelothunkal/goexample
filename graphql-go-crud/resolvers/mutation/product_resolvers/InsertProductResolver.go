package product_resolvers

import (
	"math/rand"
	"time"
	
	"github.com/graphql-go/graphql"
	model "graphql-go-crud/data/model/productDataModel"
	"graphql-go-crud/data/product_data"
)

func InsertProductResolver(params graphql.ResolveParams) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	product := model.Product{
		ID:    int64(rand.Intn(100000)), // generate random ID
		Name:  params.Args["name"].(string),
		Info:  params.Args["info"].(string),
		Price: params.Args["price"].(float64),
	}
	product_data.Products = append(product_data.Products, product)
	return product, nil
}