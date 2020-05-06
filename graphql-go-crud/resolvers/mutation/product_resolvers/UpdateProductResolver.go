package product_resolvers

import (
	"github.com/graphql-go/graphql"
	model "graphql-go-crud/data/model/productDataModel"
	"graphql-go-crud/data/product_data"
)

func UpdateProductResolver(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(int)
	name, nameOk := params.Args["name"].(string)
	info, infoOk := params.Args["info"].(string)
	price, priceOk := params.Args["price"].(float64)
	product := model.Product{}
	for i, p := range product_data.Products {
		if int64(id) == p.ID {
			if nameOk {
				product_data.Products[i].Name = name
			}
			if infoOk {
				product_data.Products[i].Info = info
			}
			if priceOk {
				product_data.Products[i].Price = price
			}
			product = product_data.Products[i]
			break
		}
	}
	return product, nil
}