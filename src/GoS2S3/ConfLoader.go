package main

import (
	"GoS2S3/salesforceUtil"
	"encoding/json"
	"log"
	"os"
)

func loadSalesforceConfigurationFromFile(configFile *SalesforceConfiguration, salesforceConnection *salesforceUtil.SF_connection) {
	salesforceConnection.TargetURI = configFile.TargetURI
	salesforceConnection.Username = configFile.Username
	salesforceConnection.Password = configFile.Password
	salesforceConnection.SecurityToken = configFile.SecurityToken
	salesforceConnection.ClientSecret = configFile.ClientSecret
	salesforceConnection.ClientId = configFile.ClientId
}

func loadAWSConfigurationFromFile(configFile *AWSConfiguration) {
	os.Setenv("AWS_PROFILE", configFile.Profile)
	os.Setenv("AWS_REGION", configFile.Region)
	os.Setenv("AWS_ACCESS_KEY_ID", configFile.Access_key_ID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", configFile.Secret_access_key)
	os.Setenv("AWS_SESSION_TOKEN", configFile.Session_token)
}

func (connectionsConf *Configuration) LoadConfigFrom(configFileName string) {
	// reading from configuration file
	configFile, _ := os.Open(configFileName) //("application-config.json")
	defer configFile.Close()
	confDecoder := json.NewDecoder(configFile)
	//configuration := Configuration{}
	err := confDecoder.Decode(&connectionsConf)
	if err != nil {
		log.Println("error:", err)
	}
}

type SalesforceConfiguration struct {
	TargetURI     string
	Username      string
	Password      string
	SecurityToken string
	ClientSecret  string
	ClientId      string
}

type AWSConfiguration struct {
	Instance_url          string
	Username              string
	Access_key_ID         string `json:"Access_key_ID"`
	Secret_access_key     string `json:"Secret_access_key"`
	Session_token         string `json:"Session_token"`
	Profile               string `json:"Profile"`
	Region                string `json:"Region"`
	S3_destination_bucket string `json:"s3_destination_bucket"`
	S3_destination_path   string `json:"s3_destination_path"`
	S3_destination_prefix string `json:"s3_destination_prefix"`
}

type Configuration struct {
	Salesforce SalesforceConfiguration `json:"Salesforce"`
	Amazon     AWSConfiguration        `json:"AWS"`
}
