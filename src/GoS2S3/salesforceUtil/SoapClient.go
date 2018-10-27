package salesforceUtil

import (
	"log"
	"net/http"
	"os"
	"strings"
	"encoding/json"
	"io/ioutil"
	"GoS2S3/SalesforceWSDL"
)

func (connection *SF_connection) GetSoapEndpoint() string {
	client := &http.Client{}
	targetInstance := connection.AuthenticationToken.Id // contains the url you have to query to get the SOAP endpoint
	
	request, _ := http.NewRequest("POST", targetInstance, nil)
	request.Header.Add("Authorization", connection.AuthenticationToken.Token_type + " " + connection.AuthenticationToken.Access_token)

	response, responseError := client.Do(request)
	if responseError != nil {
		response.Body.Close()
		log.Println("ERROR: Failed retrieving SOAP endpoint!")
		if connection.Debug {
			log.Println(responseError)
		} else {
			log.Println("Error: an error occurred during the execution, please run with -debug for more information")
		}
		
		os.Exit(1)
	}
	defer response.Body.Close() 
	
	body, _ := ioutil.ReadAll(response.Body)
	// fmt.Printf("%s\n", string(body))
	// fmt.Println(string(body))
	var jsonBody map[string] interface {}
	if err := json.Unmarshal(body, &jsonBody); err != nil {
        panic(err)
	}
	connection.OrganizationId = jsonBody["organization_id"].(string)	
	soapEndpoint := (jsonBody["urls"].(map[string] interface {}))["enterprise"]

	return soapEndpoint.(string)
}

func (connection *SF_connection) AuthenticateThroughSOAP() (SalesforceWSDL.Soap, SalesforceWSDL.BasicAuth) {
	var SF_Soap SalesforceWSDL.Soap
	var SF_BasicAuth SalesforceWSDL.BasicAuth
	
	SF_BasicAuth.Login = connection.Username
	SF_BasicAuth.Password = connection.Password
	
	connection.SoapEndpoint = connection.GetSoapEndpoint()
	connection.SoapEndpoint = strings.Replace(connection.SoapEndpoint, "{version}", "v43.0", 1)
		
	SF_Soap = *SalesforceWSDL.NewSoap(connection.SoapEndpoint, false, &SF_BasicAuth)

	var loginAttempt SalesforceWSDL.Login
	loginAttempt.Username = connection.Username
	loginAttempt.Password = connection.Password + connection.SecurityToken

	log.Println("")
	log.Printf("Logging in through SOAP ....")
	log.Println("")

	loginResponse, loginError := SF_Soap.Login(&loginAttempt)
	if(loginError != nil) {
		//fmt.Printf(" [FAILED]\n")
		log.Println("ERROR during Salesforce SOAP login")
		if connection.Debug {
			log.Println(loginError)
		}
	} else {
		//fmt.Printf(" [SUCCESS]\n")
		log.Println("LOGIN SUCCESSFUL!")
		if connection.Debug {
			log.Println("LOGIN RESPONSE: ")
			log.Println(loginResponse)
		}
		connection.SoapLogin = *loginResponse.Result		
	}
	connection.ConnectionCookies["oid"] = connection.OrganizationId
	connection.ConnectionCookies["sid"] = loginResponse.Result.SessionId
	return SF_Soap, SF_BasicAuth
}