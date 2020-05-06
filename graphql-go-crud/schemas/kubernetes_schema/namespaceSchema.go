package kubernetes_schema

import (
	"github.com/graphql-go/graphql"
)

var ListNameSpace = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NameSpaceList",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"creationTimestamp": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var CreateNameSpace = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CreateNamespace",
		Fields: graphql.Fields{
			"statusCode": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"message": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"creationTimestamp": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"resourceVersion": &graphql.Field{
				Type: graphql.String,
			},
			"selfLink": &graphql.Field{
				Type: graphql.String,
			},
			"uid": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DeleteNameSpace = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NameSpaceList",
		Fields: graphql.Fields{
			"statusCode": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"message": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"deletionTimestamp": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
