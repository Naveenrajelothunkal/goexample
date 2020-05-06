Golang-Graphql server with graphql-go/graphql

Installation : 

	To clone this Repository to your local environment, use the following command:
    
		git clone https://innersource.accenture.com/scm/ethan/golangserverapi.git

To start the app :

	•	Run the command "go run server.go" from the root folder.

Tool used to test the APIS :

	•	Postman

Usage :

	•	Start the application using the command "go run server.go".
    

	•	Generate the token - 
    
		URL :- http://localhost:8080/token 
        
		request body - {"username":"cm9vdA==","password":"cm9vdDEyMw=="}

	
    •	Execute the GraphQl Query – 
    
		URL :- http://localhost:8080/goengine?query={list{id,name,info,price}} 
        
		request body - {"token":"token-generated-from-above-step"}

Adding New API to the boiler plate :

	•	Write a graphql schema under schemas folder 
                     sample schema – 
					
                    var ProductType = graphql.NewObject(
    					graphql.ObjectConfig{
        					Name: "Product",
        					Fields: graphql.Fields{
            					"id": &graphql.Field{
                					Type: graphql.Int,
            					},
            					"name": &graphql.Field{
                					Type: graphql.String,
            					},
            					"info": &graphql.Field{
                					Type: graphql.String,
            					},
            					"price": &graphql.Field{
                					Type: graphql.Float,
            					},
        					},
    					},
					)


	•	Write a graphql resolver under resolvers folder 
    
		sample resolver – 
        
		func ProductList(params graphql.ResolveParams) (interface{}, error) {
    		return product_data.Products, nil
		}

	•	Register the Query/Mutation under query/mutaion folder respectively 
    
		sample query entry – 
        
		"list": &graphql.Field{
       		Type: graphql.NewList(product_schema.ProductType), //schema
       		Description: "Get product list",
       		Resolve: product_resolvers.ProductList,  //resolver
     	},

Naming standards which is used :

	•	Names are as important in Go as in any other language. They even have semantic effect: 		the visibility of a name outside a package is determined by whether its first character 	is upper case. It's therefore worth spending a little time talking about naming 			conventions in Go programs.
	
	•	Package Names : 
		
		•   A package in Go is simply a directory/folder with one or more .go files inside of it.
		•   package name should be good: short, concise, evocative.
		•   packages are given lower case, single-word names; there should be no need for			underscores or mixedCaps
		•   The package name is only the default name for imports; it need not be unique across 	all source code, and in the rare case of a collision the importing package can 			choose a different name to use locally
		•	Another convention is that the package name is the base name of its source 				directory; the package in src/encoding/base64 is imported as "encoding/base64" but 		has name base64, not encoding_base64 and not encodingBase64.
	
	•	File Names :

		•	Keep file and folder names short, but meaningful. Avoid unnecessary repetition and 		redundant words in file names and file paths.
		•	Generally, file names follow the same convention as package names.
		•	File names that begin with "." or "_" are ignored by the go tool
		•	Files with the suffix _test.go are only compiled and run by the go test tool.

	•	Interface Names : 

		•	By convention, one-method interfaces are named by the method name plus an -er suffix 	 or similar modification to construct an agent noun: Reader, Writer, Formatter, 		CloseNotifier etc.
		•	There are a number of such names and it's productive to honor them and the function 	names they capture. Read, Write, Close, Flush, String and so on have canonical 			signatures and meanings. To avoid confusion, don't give your method one of those 		names unless it has the same signature and meaning

	•	MixedCaps / mixedCaps : 

		•	Finally, the convention in Go is to use MixedCaps or mixedCaps rather than 				underscores to write multiword names.
		•	Exported variable/function name ==> MixedCaps 
		•	Unexported variable/function name ==> mixedCaps 


Contributing :

	•	Pull requests are welcome. Please make sure to update tests as appropriate.


