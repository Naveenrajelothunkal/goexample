package myresolvers

import (

	"graphql-go-crud/data/mydata"

	"github.com/graphql-go/graphql"
)

func MyList(params graphql.ResolveParams) (interface{}, error) {
	return mydata.MyData, nil
}
