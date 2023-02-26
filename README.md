# Golang Skeleton With Fully Managed Versions For Kick Start GoLang Project Development
Cloned and modified from https://github.com/Mindinventory/Golang-Project-Structure

It is a fully managed repository, where one can find all required components in a single package including versioning for REST APIs and you do not need to set up each time they start with any crucial work in Golang.


## Prerequisite

One need to install the latest version of Go i.e 1.12 (Released in Feb 2019) from https://golang.org/dl/ and setup GOROOT and GOPATH.

## Components 
<center><img src="https://raw.githubusercontent.com/Mindinventory/mklogistics/master/gif.gif"></center>


### 1. ApiHelpers
Basically contains the helper functions used in returning api responses, HTTP status codes, default messages etc.

### 2. Controllers
Contains handler functions for particular route to be called when an api is called.

### 3. Helpers
Contains helper functions used in all apis

### 4. Middlewares
Middleware to be used for the project

### 5. Models
Database tables to be used as models struct

### 6. Resources
Resources contains all structures other than models which can be used as responses

### 7. Routers
Resources define the routes for your project

### 8. Seeder
It is optional, but if you want to insert lots of dummy records in your database, then you can use seeder.

### 9. Services
All the core apis for your projects should be within services.

### 10. Storage
It is generally for storage purpose.

### 11. Templates
Contains the HTML templates used in your project

### 12. .env
Contains environment variables.


## Steps to Follow

. Find and replace all instances of "mklogistics" with your app name.

. For running the server you have to run following command in the terminal.
        ```go run main.go
        ```
  It will start your server at the port you have mentioned in the ```.env``` file.
  
. To run the server in a port other than the default, run following command.
        ```go run main.go <specific port>```
        
. To create a build for your project and uploaded in the server, one need to run following command.
        ```go build```
        
       
## API with versioning

# For using version 1 api
```127.0.0.1:8099/api/v1/user-list```

# For using version 2 api
```127.0.0.1:8099/api/v2/user-list```
