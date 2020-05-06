package mongo_schema

import (
	"github.com/graphql-go/graphql"
)

var ToolType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tool",
		Fields: graphql.Fields{
			"Order_Id": &graphql.Field{
				Type: graphql.Int,
			},
			"Created_On": &graphql.Field{
				Type: graphql.String,
			},
			"Updated_On": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)