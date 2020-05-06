package kubernetes

import (
	"io/ioutil"
	"encoding/json"
	"log"

	"github.com/graphql-go/graphql"

	"graphql-go-crud/data/connector/k8sConnector"
)

type deleteNameSpaceResponse struct {
	StatusCode int
	Message string
	DeletionTimestamp string
	Name string
}

func DeleteNameSpace(params graphql.ResolveParams) (interface{}, error) {
	var response deleteNameSpaceResponse

	nameSpace,ok := params.Args["nameSpace"].(string)
	if !ok {return nil,nil}
	clusterId,ok := params.Args["clusterId"].(string)
	if !ok {return nil,nil}

	client,req,err := k8sConnector.ConnectCluster("DELETE","/api/v1/namespaces/"+nameSpace,clusterId,nil)
	if err != nil {
		log.Print(err)
		return deleteNameSpaceResponse{StatusCode:1},err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return nil,nil
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatal(err)
		return nil,err
	}

	if result["kind"] == "Namespace" {
		metadata := result["metadata"].(map[string]interface{})
		response = deleteNameSpaceResponse{
			StatusCode:0,
			Message:"Successfully deleted namespace",
			DeletionTimestamp:metadata["deletionTimestamp"].(string),
			Name:metadata["name"].(string),
		}
		return response,nil
	} else if result["kind"] == "Status" {
		response = deleteNameSpaceResponse{
			StatusCode:1,
			Message:result["message"].(string),
		}
		return response,nil
	}
	
	return nil,nil
}
