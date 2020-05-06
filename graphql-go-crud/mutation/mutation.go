package mutation

import (
	"github.com/graphql-go/graphql"

	//schemas
	"graphql-go-crud/schemas/kubernetes_schema"
	"graphql-go-crud/schemas/mongo_schema"
	"graphql-go-crud/schemas/product_schema"
	"graphql-go-crud/schemas/user_schema"

	//resolvers
	"graphql-go-crud/resolvers/mutation/mongoResolvers"
	"graphql-go-crud/resolvers/mutation/product_resolvers"
	"graphql-go-crud/resolvers/mutation/updateskillresolvers"
	"graphql-go-crud/resolvers/mutation/userinfoResolvers"
	"graphql-go-crud/resolvers/query_resolvers/kubernetes"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/* Create new product item
		http://localhost:8080/goengine?query=mutation+_{insertProduct(name:"Inca Kola",info:"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)",price:1.99){id,name,info,price}}
		*/
		"insertProduct": &graphql.Field{
			Type:        product_schema.ProductType,
			Description: "Create new product",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"info": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
			},
			Resolve: product_resolvers.InsertProductResolver,
		},

		/* Creat a new skill
		http://localhost:8080/goengine?query=mutation+_{insertSkill(skill_id:1,name:"Golang",rating:8){skill_id, name, rating}}
		*/
		"insertSkill": &graphql.Field{
			Type:        mongo_schema.SkillType,
			Description: "Create new skill",
			Args: graphql.FieldConfigArgument{
				"skill_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"rating": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: mongoResolvers.InsertSkill,
		},
		/* Update product by id
		   http://localhost:8080/goengine?query=mutation+_{updateProduct(id:1,price:3.95){id,name,info,price}}
		*/
		"updateProduct": &graphql.Field{
			Type:        product_schema.ProductType,
			Description: "Update product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"info": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
			},
			Resolve: product_resolvers.UpdateProductResolver,
		},
		/* Update skill
		   http://localhost:8080/goengine?query=mutation+_{updateskill(skill_id:1,name:"Mongodb",rating:5){skill_id,name,rating}}
		*/
		"updateskill": &graphql.Field{
			Type:        mongo_schema.SkillType,
			Description: "Update skill data",
			Args: graphql.FieldConfigArgument{
				"skill_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"rating": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: updateskillresolvers.UpdateSkillResolver,
		},
		/* Delete product by id
		   http://localhost:8080/goengine?query=mutation+_{deleteProduct(id:1){id,name,info,price}}
		*/
		"deleteProduct": &graphql.Field{
			Type:        product_schema.ProductType,
			Description: "Delete product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: product_resolvers.DeleteProductResolver,
		},

		/* Insert to Mongo
		   http://localhost:8080/goengine?query=mutation+_{insertMongo(Order_Id:"1234",Created_On:"feb",Updated_On:"xyz"){Order_Id,Created_On,Updated_On}}
		*/
		"insertMongo": &graphql.Field{
			Type:        mongo_schema.ToolType,
			Description: "Insert to mongo",
			Args: graphql.FieldConfigArgument{
				"Order_Id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"Created_On": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"Updated_On": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: mongoResolvers.InsertOneMongoDb,
		},

		/* Insert to Mongo
		   http://localhost:8080/goengine?query=mutation+_{insertUser(id:1,Name:"Raj",Place:"Bangalore",Created_On:"feb",Updated_On:"mar"){id,Name,Place,Created_On,Updated_On}}
		*/
		"insertUser": &graphql.Field{
			Type:        user_schema.UserInfo,
			Description: "Insert User info to mongo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"Name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"Place": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"Created_On": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"Updated_On": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: userinfoResolvers.UserinfoMongoDb,
		},

		/* http://localhost:8080/goengine?query=mutation+_{createNamespace(clusterId:"devplatform",nameSpace:"testing"){statusCode,message,name}} */
		"createNamespace": &graphql.Field{
			Type:        kubernetes_schema.CreateNameSpace,
			Description: "kubernetes cluster create namespace",
			Args: graphql.FieldConfigArgument{
				"clusterId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"nameSpace": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: kubernetes.CreateNameSpace,
		},

		/* http://localhost:8080/goengine?query=mutation+_{deleteNamespace(clusterId:"devplatform",nameSpace:"testing1"){Name,statusCode,deletionTimestamp,message}} */

		"deleteNamespace": &graphql.Field{
			Type:        kubernetes_schema.DeleteNameSpace,
			Description: "kubernetes cluster delete namespace",
			Args: graphql.FieldConfigArgument{
				"clusterId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"nameSpace": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: kubernetes.DeleteNameSpace,
		},
	},
})
