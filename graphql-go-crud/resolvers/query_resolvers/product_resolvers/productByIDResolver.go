package product_resolvers

import (
	"github.com/graphql-go/graphql"
	"graphql-go-crud/data/product_data"
)

func ProductByIDResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if ok {
		// Find product
		for _, product := range product_data.Products {
			if int(product.ID) == id {
				return product, nil
			}
		}
	}
	return nil, nil
}