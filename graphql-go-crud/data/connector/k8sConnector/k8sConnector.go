package k8sConnector

import (
    "log"
	"net/http"
	"crypto/tls"
	"bytes"
	"errors"

	"graphql-go-crud/config"
)

func ConnectCluster(requestMethod,apiPath,clusterID string, requestBody []byte) (*http.Client,*http.Request,error) {
	var clusterUrl string
	var clusterToken string

	for _,res := range config.Config.Cluster {
		if res.Id == clusterID {
			clusterUrl = res.Url
			clusterToken = res.Token
		}
	}

	if clusterUrl == "" || clusterToken == "" {return nil,nil,errors.New("Cluster not found")}

	bearer := "Bearer " + clusterToken
	req, err := http.NewRequest(requestMethod, clusterUrl+apiPath, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Print(err)
		return nil,nil,err
	}
	req.Header.Add("Authorization", bearer)
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	client := &http.Client{Transport: tr}

	return client,req,nil
}