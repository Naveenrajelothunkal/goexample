package mongo_schema

import (
	"github.com/graphql-go/graphql"
)

var SkillType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Skill",
		Fields: graphql.Fields{
			"skill_id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"rating": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
