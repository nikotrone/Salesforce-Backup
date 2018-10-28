# Go Salesforce Backup
The Salesforce platform offers the possibility to request a backup of the entire database. When it is ready, the files (512MB each) are listed for download.

This tool scraps the page and downloads the files one by one and pushes them to an S3 bucket used to store the backup.

## Setting up environment
To install the Golang environment visit https://golang.org/doc/install

On the command line write "go get golang.org/x/net/html" to fetch the HTML tokenizer dependancy

On the command line write "go get -u github.com/aws/aws-sdk-go" to fetch the AWS SDK for Golang

More information at https://golangbot.com/golang-tutorial-part-1-introduction-and-installation/
Tool used to generate WSDL definition at [Hooklift Github repository](https://github.com/hooklift/gowsdl)

## Program Installation

Open the terminal and go in your GOROOT directory, ie:
```console
# cd ~/go/src
```
Clone the project from github with:
```console
# git clone https://github.com/nikotrone/go-Salesforce-backup.git
```

Inside the program folder there is an example configuration file called *application-config.example.json*

Clone it to *application-config.json* and edit to add your personal configurations

## Run it in Docker

The application has already a basic Dockerfile and a docker-compose.yml so it can be run in a docker container with a different installation process:
- by simply cloning the project, setting up your *application-config.json* and running this commands on a shell
```console
# docker-compose up --build
```

In alternative you can use the basic docker commands to build and run the image/container
```console
# docker build -t imagename:tag .
```
When done you will have an image called gos2s3 in your local repository plus an untagged image used to build the executable

## External references:
Program used to generate WSDL https://github.com/hooklift/gowsdl
### Hacktoberfest
    +[Link cek progres ](https://hacktoberfest.digitalocean.com)

