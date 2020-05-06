package user_schema

import (
	"github.com/graphql-go/graphql"
)

var UserInfo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Place": &graphql.Field{
				Type: graphql.String,
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
