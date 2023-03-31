# Go-Api-Gin-Gorm
REST API made with GO-GIN-GORM
This API only watches for :PORT/posts (CRUD) requests
Get (:PORT/posts) => returns all posts from db
Post (:PORT/posts) => create and return a posts by sendind {"Title":"some","Body":"some-body"}

Also includes patch/put/get:id routes... just check yourself.
### Adding new packages
 go get gorm.io/driver/mysql

### I'm using GORM ORM with PostgresQL, to enable it fill the DB_URI in .env like:
POSTGRES_URI="host=localhost user=ian password=1708 dbname=dianaOnRails_development port=5432 sslmode=disable"

### :PORT -- Gin uses :8080 as default, if u want to use other just set in .env like:
PORT = 3000 

### Running migrations:
go run migrate/migrate.go

## Running aplication:
go run main.go 
OR (the best like nodemon)
CompileDaemon -command="./goagain"
***CompileDaemon watches the changes and restart, no need to break and start again.



See your, yours Ian