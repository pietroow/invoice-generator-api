- go get github.com/githubnemo/CompileDaemon
- go install github.com/githubnemo/CompileDaemon

- go get github.com/joho/godotenv # load environment variables
- go get -u github.com/gin-gonic/gin #web framework
- go get -u gorm.io/gorm # widely used for mapping orm
- go get -u gorm.io/driver/postgres # postgres driver

Project must be on 
${GOPATH}/src

To download dependencies we run
$ go mod download

To run the migrations run
$ go run migrate/migrate.go

Run the project in live reload mode
$ CompileDaemon -command="./invoice-generator-api"