package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_"github.com/lib/pq"
)

func main() {
	r := gin.Default()
	r.GET("/ping", )
	r.POST("/users", create)
	r.Run(":50051")
}

func PingHandler(c *gin.Context){
	res := Response{Message: "pong"}
	c.JSON(http.StatusOK, res)
}

// var db *sql.DB

func initDb()(*sql.DB, error){
	db, err := sql.Open("postgres", "postgres://postgres:Aa@@2004@postgres-db:5432/postgres?sslmode=disable")
	return db, err
}

type Response struct{
	Message string `json:"message"`
}

type User struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func create(c *gin.Context){
	db, err := initDb()
	if err != nil {
		return
	}
	defer db.Close()
	var req User
	err = c.ShouldBindJSON(&req)
	if err != nil {
		println(err)
		return
	}
	query := "insert into usersm(id, name, email) values($1, $2, $3)"
	_, err = db.Exec(query, req.ID, req.Name, req.Email)
	if err != nil {
		return 
	}
	println("username: ", req.Name, "\nEmail: ", req.Email)
}

