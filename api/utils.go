package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = []byte("secret -key")
)

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

func FindDatabase(tag, password string) bool {
	CredDB, err := sql.Open("sqlite3", "models/CredDB.db")
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}
	defer CredDB.Close()

	var dbID int
	err = CredDB.QueryRow(`SELECT DBID FROM DatabaseCred WHERE DBTag=? AND DBPassword=?`, tag, password).Scan(&dbID)
	return err == nil
}

func UserExists(username, password string) bool {
	UserDB, err := sql.Open("sqlite3", "models/UserDB.db")
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}
	defer UserDB.Close()
	var uname string
	err = UserDB.QueryRow(
		`SELECT Username 
		 FROM Users 
		 WHERE Username=? AND Password=?`,
		username, password,
	).Scan(&uname)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	fmt.Println("Found username:", uname)

	return uname != ""

}

func CreateToken(username, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"password": password,
			"exp":      time.Now().Add(time.Minute * 2).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
