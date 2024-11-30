package main

import (
	"database/sql"
	"log"
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

func UserExists(username, password, dbname, dbpassword string) bool {
	UserDB, err := sql.Open("sqlite3", "models/UserDB.db")
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}
	defer UserDB.Close()
	var id int
	err = UserDB.QueryRow(`select Username,Password,DBTag,DBPassword from Users where Username=? and DBTag=? and Password=? and DBPassword=?`, username, dbname, password, dbpassword).Scan(&id)

	return err == nil

}
