Golang API Echo Framework MongoBD 

how to run this project in local
 
and I am trying to run it on my local machine but unable system?
Make sure you have enabled Go modules by running the command export GO111MODULE=on in your terminal.
Initialize a new Go module by running the command go mod init <module-name> in the directory where your main.go file is located. Replace <module-name> with a name of your choice.
Download the required dependencies by running the command go get -u github.com/labstack/echo/v4 github.com/labstack/echo-contrib/tree/master/mongoDB.
Run your main.go file using the command go run main.go.
If you still encounter issues, you can try deleting your Go module cache by running the command go clean -modcache and then repeating the above steps.

Additionally, make sure that you have MongoDB installed and running on your local machine. You can download MongoDB from the official website and follow the installation instructions for your operating system. Once installed, you can start the MongoDB main by running the command mongod in your terminal.

I hope this helps! Let me know if you have any further questions.

Complete data in  .ENV 

connect you MONGO URL  :"DOCKER " mongodb://localhost:27017/<DATABASE>?ssl=false&auth
or :  mongodb+srv://USER:PASSWORD@YOUR-CLUSTER.mongodb.net/   

OR for connection check  /database/connect.go 

### Installation and Setup ###

to start server  cd  youruser/go-echo/go run main.go

you can use Postman or Insomnia to test routes

## with Docker ##
open your terminal  cd/project folder  and start: docker compose up

then open postman or insomnia  and go to  http://localhost:8085/api/v1/ "all routes""

Thanks 

## PUBLIC ##
the ./public folder is only as a example feel free to  change it for you own template
plublic view work in:   http://localhost:8085/
