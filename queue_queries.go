package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB_NAME string
var DB_HOST string
var DB_USERNAME string
var DB_PASSWORD string
var dbURL string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	dbURL = DB_USERNAME + ":" + DB_PASSWORD + DB_HOST + "/" + DB_NAME
}

func addQueue(q Queue) Queue {
	db, err := sql.Open("mysql", dbURL)
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO queue VALUES( ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() //
	if err != nil {
		fmt.Println("Failed to connect", err)
		return nil
	}
}

func getQueues() (queues []Queue) {
	db, err := sql.Open("mysql", dbURL)
	defer db.Close()
	if err != nil {
		fmt.Println("Failed to connect", err)
		return
	}

	rows, err := db.Query("SELECT * FROM queue")
	if err != nil {
		fmt.Println("Failed to run query", err)
		return
	}

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return
	}

	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		id, _ := strconv.Atoi(result[0])
		queue := Queue{
			id,
			result[1],
			result[2],
			result[3],
			stringToBool(result[4]),
			stringToBool(result[5]),
			stringToBool(result[6]),
			stringToBool(result[7]),
			stringToBool(result[8]),
		}
		queues = append(queues, queue)
	}
	return
}

func stringToBool(s string) bool {
	if s == "0" || s == "" {
		return false
	} else {
		return true
	}
}
