package product_resolvers

import (
	"github.com/graphql-go/graphql"
	model "graphql-go-crud/data/model/productDataModel"
	"graphql-go-crud/data/product_data"
)

func DeleteProductResolver(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(int)
	product := model.Product{}
	for i, p := range product_data.Products {
		if int64(id) == p.ID {
			product = product_data.Products[i]
			// Remove from product list
			product_data.Products = append(product_data.Products[:i], product_data.Products[i+1:]...)
		}
	}

	return product, nil
}