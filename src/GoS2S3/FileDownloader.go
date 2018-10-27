package main

import(
	"log"
	"net/http"
	"os"
	"io"	
)

func downloadFileFromUrl (targetFileUrl string, parameters map[string] interface {}) (fileName string, someError error) {
	client := &http.Client{}

	request, requestError := http.NewRequest("GET", targetFileUrl, nil)
	if requestError != nil {
		log.Printf("Error creating the download request for %s: \n\t - %s", targetFileUrl, requestError)
		someError = requestError
		return
	}
	
	fileName = request.URL.Query()["fileName"][0]
	
	// Use array instead of appending one by one
	sliceOfCookies := createCookieList(parameters)
	for _, value := range sliceOfCookies {
		request.AddCookie(&value)
	}
	

	// Create the file
	downloadedFile, fileCreationErr := os.Create("tmp/" + fileName)
    if fileCreationErr != nil {
        log.Printf("Error creating the output file %s: \n\t - %s", fileName, fileCreationErr)				
		someError = fileCreationErr
		return
    }
	defer downloadedFile.Close()
	

	// Do the request
	response, responseError := client.Do(request)
	if responseError != nil {
		log.Printf("Error doing the downloading file %s: \n\t - %s", fileName, responseError)		
		someError = responseError
		return		
	}
	defer response.Body.Close()
	

	// Write the body to file
	_, savingError := io.Copy(downloadedFile, response.Body)
	if savingError != nil {
		log.Printf("Error while saving the file %s: \n\t - %s", fileName, savingError)				
		someError = savingError		
		return
	}

	return
}

func createCookieList (inputMap map[string] interface {}) (sliceOfCookies []http.Cookie) {
	for key, _ := range inputMap {
		sliceOfCookies = append(sliceOfCookies, http.Cookie{ Name: key, Value: inputMap[key].(string)})
	}
	return 
}
