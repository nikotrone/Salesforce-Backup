package main

import (	
	"os"
	"log"	
	"strconv"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func uploadFileToS3(currentSession *session.Session, applicationConfiguration AWSConfiguration, filename string) (transferResult string, transferError error) {
			
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(currentSession)

	fileReader, readingError  := os.Open("tmp/" + filename)
	if readingError != nil {
		log.Printf("Failed to open file %q, %v", filename, readingError)
		return "Failed to open file " + filename, readingError
	}
	defer fileReader.Close()

	// Upload the file to S3.
	uploadResult, uploadError := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(applicationConfiguration.S3_destination_bucket),
		Key:    aws.String(applicationConfiguration.S3_destination_path + strconv.FormatInt(todayEpoch, 10) + "/" + applicationConfiguration.S3_destination_prefix + filename),
		Body:   fileReader,
	})
	if uploadError != nil {
		log.Printf("Failed to upload file, %v", uploadError)
		return "Failed to upload file " + filename, uploadError
	}
	log.Printf("File uploaded to, %s\n", aws.StringValue(&uploadResult.Location))

	os.Remove("tmp/" + filename)

	return "Success", nil
}