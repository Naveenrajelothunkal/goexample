package queries

import (
	"github.com/graphql-go/graphql"

	//schemas
	"graphql-go-crud/schemas/kubernetes_schema"
	"graphql-go-crud/schemas/mongo_schema"
	"graphql-go-crud/schemas/myschema"
	"graphql-go-crud/schemas/product_schema"
	"graphql-go-crud/schemas/user_schema"

	//resolvers
	"graphql-go-crud/resolvers/query_resolvers/kubernetes"
	"graphql-go-crud/resolvers/query_resolvers/myresolvers"
	"graphql-go-crud/resolvers/query_resolvers/product_resolvers"
	"graphql-go-crud/resolvers/query_resolvers/skillresolvers"
	"graphql-go-crud/resolvers/query_resolvers/user_resolvers"
	
)

var Query = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single product by id
			   http://localhost:8080/goengine?query={product(id:1){name,info,price}}
			*/
			"product": &graphql.Field{
				Type:        product_schema.ProductType,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: product_resolvers.ProductByIDResolver,
			},
			/* Get (read) product list
			   http://localhost:8080/goengine?query={list{id,name,info,price}}
			*/
			"list": &graphql.Field{
				Type:        graphql.NewList(product_schema.ProductType),
				Description: "Get product list",
				Resolve:     product_resolvers.ProductList,
			},
			/* Get (read) mylist list
			   http://localhost:8080/goengine?query={mylist{id,name,info,price}}
			*/
			"mylist": &graphql.Field{
				Type:        graphql.NewList(myschema.MyType),
				Description: "Get my list",
				Resolve:     myresolvers.MyList,
			},
			/* Get (read) user list
			   http://localhost:8080/goengine?query={getUserList{id,Name,Place,Created_On,Updated_On}}
			*/
			"getUserList": &graphql.Field{
				Type:        graphql.NewList(user_schema.UserInfo),
				Description: "Get user list",
				Resolve:     user_resolvers.UserListResolver,
			},

			/* Get skill list
			http://localhost:8080/goengine?query={skilllist{ skill_id, name, rating }}
			*/

			"skilllist": &graphql.Field{
				Type:        graphql.NewList(mongo_schema.SkillType),
				Description: "Get skill list",
				Resolve:     skillresolvers.SkillListResolver,
			},

			// http://localhost:8080/goengine?query={namespaceList(clusterId:"devplatform"){name,creationTimestamp}}
			"namespaceList": &graphql.Field{
				Type:        graphql.NewList(kubernetes_schema.ListNameSpace),
				Description: "kubernetes cluster namespace list",
				Args: graphql.FieldConfigArgument{
					"clusterId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: kubernetes.ListNameSpace,
			},
		},
	})
