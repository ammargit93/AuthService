package main

import (
	"database/sql"
	"fmt"
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

func Register(ctx *gin.Context) {
	var user User

	UserDB, err := sql.Open("sqlite3", "models/UserDB.db")
	if err != nil {
		panic(err)
	}
	defer UserDB.Close()

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind json"})
	}
	var flag = FindDatabase(user.DBTag, user.DBPassword)
	if !flag {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Make a database first you can do it using /assign-dbcred"})
		return
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Successfully registered!"})
	}
	if UserExists(user.Username, user.Password, user.DBTag, user.DBPassword) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User Already exists"})
		return
	}
	_, err = UserDB.Exec(`insert into Users (Username, Password, DBTag, DBPassword) values (?,?,?,?)`, user.Username, user.Password, user.DBTag, user.DBPassword)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Username , password, DBname, DBpassword must all be unique"})
		panic(err)
	}
	ctx.JSON(200, gin.H{"message": "Inserted Successfully"})
}

func AssignCredentials(ctx *gin.Context) {
	var cred DBCred

	CredDB, err := sql.Open("sqlite3", "models/CredDB.db")
	if err != nil {
		panic(err)
	}
	defer CredDB.Close()

	if err := ctx.BindJSON(&cred); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind json to cred"})
	}

	_, err = CredDB.Exec(`insert into DatabaseCred (DBTag, DBPassword) values (?,?)`, cred.DBTag, cred.DBPassword)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "DBname, DBpassword must all be unique"})
		panic(err)
	}
	ctx.JSON(200, gin.H{
		"message": "Credentials Inserted Successfully",
	})
}

func ListRecords(ctx *gin.Context) {
	dbTag := ctx.Param("list")
	var user User

	UserDB, err := sql.Open("sqlite3", "models/UserDB.db")
	if err != nil {
		panic(err)
	}
	defer UserDB.Close()

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind json"})
	}

	row, err := UserDB.Query(`select Username, Password, DBTag, DBPassword from Users where DBTag=?`, dbTag)
	if err != nil {
		panic(err)
	}

	var UserArr []map[string]interface{}
	for row.Next() {
		var username, password, dbtag, dbpassword string
		row.Scan(&username, &password, &dbtag, &dbpassword)
		fmt.Println(username, password, dbtag, dbpassword)
		user := map[string]interface{}{
			"Username":         username,
			"Password":         password,
			"DatabaseName":     dbtag,
			"DatabasePassword": dbpassword,
		}
		UserArr = append(UserArr, user)

	}
	fmt.Println(UserArr)
	ctx.JSON(http.StatusAccepted, gin.H{"data": UserArr})

}

func Login(ctx *gin.Context) {
	var user User

	UserDB, err := sql.Open("sqlite3", "models/UserDB.db")
	if err != nil {
		panic(err)
	}
	defer UserDB.Close()

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind json"})
	}
	if UserExists(user.Username, user.Password, user.DBTag, user.DBPassword) {

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Register first using /register"})
	}
}
