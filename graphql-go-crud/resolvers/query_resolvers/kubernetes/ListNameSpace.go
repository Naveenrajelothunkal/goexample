package kubernetes

import (
	"io/ioutil"
	"log"
	"encoding/json"

	"github.com/graphql-go/graphql"

	"graphql-go-crud/data/connector/k8sConnector"
)

type listNameSpaceResponse struct {
	Name string `json:"name"`
	CreationTimestamp string `json:"creationTimestamp"`
}

func ListNameSpace(params graphql.ResolveParams) (interface{}, error) {
	clusterId,ok := params.Args["clusterId"].(string)
	if ok {
		client,req,err := k8sConnector.ConnectCluster("GET","/api/v1/namespaces",clusterId,nil)
		if err != nil {
			log.Print("res ", err)
		}
		resp, err := client.Do(req)
    	if err != nil {
        	log.Println("Error on response.\n[ERRO] -", err)
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		
		var result map[string]interface{}
		if err := json.Unmarshal([]byte(body), &result); err != nil {
			log.Fatal(err)
			return nil,err
		}
	
		var response []listNameSpaceResponse
		if result["kind"] == "NamespaceList" {
			namespaceList := result["items"].([]interface{})

			for _,res := range namespaceList {
				nameSpace := res.(map[string]interface{})
				metadata := nameSpace["metadata"].(map[string]interface{})
				val := listNameSpaceResponse{Name:metadata["name"].(string),CreationTimestamp:metadata["creationTimestamp"].(string)}
				response = append(response,val)
			}
		}
		return response,nil
	}

	return nil,nil
}