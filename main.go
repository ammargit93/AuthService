package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username string
	Password string
	DBTag    string
}

func SetupSchema() error {
	UserDB, err := sql.Open("sqlite3", "models/UserDB.db")
	if err != nil {
		panic(err)
	}
	defer UserDB.Close()

	CreateUserQuery := `create table if not exists Users(
						UID integer primary key autoincrement,
						Username text unique not null,
						Password text unique not null,
						DBtag text unique not null
						)`
	_, err = UserDB.Exec(CreateUserQuery)
	if err != nil {
		panic(err)
	}
	return nil

}

func main() {
	fmt.Println("HEllo")
	SetupSchema()

	router := gin.Default()
	router.POST("/register", func(ctx *gin.Context) {
		var user User

		UserDB, err := sql.Open("sqlite3", "models/UserDB.db")
		if err != nil {
			panic(err)
		}
		defer UserDB.Close()

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind json"})
		}
		fmt.Println(user.Password)
		fmt.Println(user.Username)
		_, err = UserDB.Exec(`insert into Users (Username, Password, DBTag) values (?,?,?)`, user.Username, user.Password, user.DBTag)

		if err != nil {
			panic(err)
		}
		ctx.JSON(200, gin.H{
			"message": "Inserted Successfully",
		})
	})

	router.Run("localhost:8080")
}
