package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username   string
	Password   string
	DBTag      string
	DBPassword string
}

type DBCred struct {
	DBTag      string
	DBPassword string
}

func SetupSchema() error {
	UserDB, err := sql.Open("sqlite3", "models/UserDB.db")
	if err != nil {
		panic(err)
	}
	CredDB, err := sql.Open("sqlite3", "models/CredDB.db")
	if err != nil {
		panic(err)
	}
	defer UserDB.Close()
	defer CredDB.Close()

	CreateUserQuery := `create table if not exists Users(
						UID integer primary key autoincrement,
						Username text unique not null,
						Password text unique not null,
						DBtag text not null,
						DBPassword text not null
						)`

	CreateDatabaseCredQuery := `create table if not exists DatabaseCred(
						DBID integer primary key autoincrement,
						DBTag text unique not null,
						DBPassword text unique not null
						)`
	_, err = UserDB.Exec(CreateUserQuery)
	if err != nil {
		panic(err)
	}
	_, err = CredDB.Exec(CreateDatabaseCredQuery)
	if err != nil {
		panic(err)
	}
	return nil

}

func FindDatabase(password, tag string) bool {
	CredDB, err := sql.Open("sqlite3", "models/CredDB.db")
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}
	defer CredDB.Close()

	_, err = CredDB.Query(`SELECT Username, Password, DBTag, DBPassword FROM DatabaseCred WHERE DBTag=? AND DBPassword=?`, tag, password)
	return err == nil
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

		var flag = FindDatabase(user.DBTag, user.Password)
		if !flag {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Make a database first you can do it using /assign-dbcred"})
			return
		}
		_, err = UserDB.Exec(`insert into Users (Username, Password, DBTag, DBPassword) values (?,?,?,?)`, user.Username, user.Password, user.DBTag, user.DBPassword)

		if err != nil {
			panic(err)
		}
		ctx.JSON(200, gin.H{
			"message": "Inserted Successfully",
		})
	})

	router.POST("/assign-dbcred", func(ctx *gin.Context) {
		var cred DBCred

		CredDB, err := sql.Open("sqlite3", "models/CredDB.db")
		if err != nil {
			panic(err)
		}
		defer CredDB.Close()

		if err := ctx.BindJSON(&cred); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind json to cred"})
		}
		fmt.Println(cred.DBTag)
		fmt.Println(cred.DBPassword)
		_, err = CredDB.Exec(`insert into DatabaseCred (DBTag, DBPassword) values (?,?)`, cred.DBTag, cred.DBPassword)

		if err != nil {
			panic(err)
		}
		ctx.JSON(200, gin.H{
			"message": "Credentials Inserted Successfully",
		})
	})

	router.Run("localhost:8080")
}
