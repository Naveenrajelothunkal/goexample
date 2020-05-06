package mydata

import (
	model "graphql-go-crud/data/model/mydatamodel"
)

var MyData = []model.Mydata{
	{
		ID:    1,
		Name:  "AWS",
		Info:  "Excellent product",
		Price: 12000,
	},
	{
		ID:    2,
		Name:  "Azure",
		Info:  "Good product",
		Price: 6000,
	},
}
