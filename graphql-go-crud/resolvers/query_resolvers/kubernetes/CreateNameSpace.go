package kubernetes

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/graphql-go/graphql"

	"graphql-go-crud/data/connector/k8sConnector"
)

type createNameSpaceResponse struct {
	StatusCode int
	Message string
	CreationTimestamp string
	Name string
	ResourceVersion string
	SelfLink string
	Uid string
}

func CreateNameSpace(params graphql.ResolveParams) (interface{}, error) {
	var response createNameSpaceResponse

	clusterId,ok := params.Args["clusterId"].(string)
	if !ok {return nil,nil}
	nameSpace,ok := params.Args["nameSpace"].(string)
	if !ok {return nil,nil}

	var requestBody = []byte(`{"apiVersion": "v1","kind": "Namespace","metadata":{"name":"` + nameSpace + `"}}`)
	
	client,req,err := k8sConnector.ConnectCluster("POST","/api/v1/namespaces",clusterId,requestBody)
	if err != nil {
		log.Print(err)
		return createNameSpaceResponse{StatusCode:1},err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return createNameSpaceResponse{StatusCode:1},err
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
		response = createNameSpaceResponse{
			StatusCode:0,
			Message:"Successfully created namespace",
			CreationTimestamp:metadata["creationTimestamp"].(string),
			Name:metadata["name"].(string),
			ResourceVersion:metadata["resourceVersion"].(string),
			SelfLink:metadata["selfLink"].(string),
			Uid:metadata["uid"].(string),
		}
		return response,nil
	} else if result["kind"] == "Status" {
		response = createNameSpaceResponse{
			StatusCode:1,
			Message:result["message"].(string),
		}
		return response,nil
	}
	return nil,nil
}
