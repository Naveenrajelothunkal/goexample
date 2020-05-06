package mongoResolvers

import (
	"graphql-go-crud/data/model/mongoModel/skilldetailsmodel"
	"graphql-go-crud/data/services/mongoDbServices"

	"github.com/graphql-go/graphql"
)

func InsertSkill(params graphql.ResolveParams) (interface{}, error) {
	data := skilldetailsmodel.SkillType{
		Skill_id: params.Args["skill_id"].(int),
		Name:     params.Args["name"].(string),
		Rating:   params.Args["rating"].(int),
	}
	mongoDbServices.InsertSkillToMongo(data)
	return data, nil
}
