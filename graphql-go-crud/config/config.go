package config

import (
	b64 "encoding/base64"
	json "encoding/json"
	"log"
	"os"
)

type clusterInfo struct {
	Id    string
	Url   string
	Token string
}

type Configuration struct {
	TokenExpiry int
	DbUrl       string
	Cluster     []clusterInfo
}

var Config Configuration

func ConfigureServer() {
	secret := b64.StdEncoding.EncodeToString([]byte("my_secret_key"))
	user := b64.StdEncoding.EncodeToString([]byte("root"))
	password := b64.StdEncoding.EncodeToString([]byte("root123"))

	os.Setenv("secretkey", secret)
	os.Setenv("serverUser", user)
	os.Setenv("serverPassword", password)

	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal(err)
	}
}
