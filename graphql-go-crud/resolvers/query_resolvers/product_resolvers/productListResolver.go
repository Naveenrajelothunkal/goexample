package product_resolvers

import (
	"github.com/graphql-go/graphql"
	"graphql-go-crud/data/product_data"
)

func ProductList(params graphql.ResolveParams) (interface{}, error) {
	return product_data.Products, nil
}