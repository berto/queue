package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
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
	dbURL = DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":3306)" + "/" + DB_NAME
}

func insertQueue(q Queue) (Queue, string) {
	db, err := sqlx.Connect("mysql", dbURL)

	if err != nil {
		fmt.Println("Failed to connect", err)
		return q, "Failed to connect"
	}
	defer db.Close()

	tx := db.MustBegin()

	result, err := tx.NamedExec(
		`INSERT INTO queue (name, location, question, googled, asked_student, has_debugged, contacted, completed) 
		VALUES (:name, :location, :question, :googled, :asked_student, :has_debugged, :contacted, :completed)`,
		q,
	)
	if err != nil {
		fmt.Println("Failed to insert", err)
		return q, "Failed to insert"
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("Failed to commit transaction", err)
		return q, "Failed to commit transaction"
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Failed to get id", err)
		return q, "Failed to get id"
	}

	q.ID = int(id)
	return q, ""
}

func completeQueue(id int) (Queue, string) {
	var queue Queue

	db, err := sqlx.Connect("mysql", dbURL)
	if err != nil {
		fmt.Println("Failed to connect", err)
		return queue, "Failed to connect"
	}
	defer db.Close()

	err = db.Get(&queue, "SELECT * FROM queue WHERE id=?", strconv.Itoa(id))
	if queue.ID != id || err != nil {
		fmt.Println("Failed to select queue", err)
		return queue, "Failed to select queue"
	}

	tx := db.MustBegin()

	result := tx.MustExec("UPDATE queue SET completed=true WHERE id=?", strconv.Itoa(id))
	rows, err := result.RowsAffected()

	if rows < 1 {
		fmt.Println("Failed to delete: record already completed")
		return queue, "Failed to delete: record already completed"
	}

	if err != nil {
		fmt.Println("Failed to delete", err)
		return queue, "Failed to delete"
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("Failed to commit transaction", err)
		return queue, "Failed to commit transaction"
	}

	queue.Completed = true
	return queue, ""
}

func getQueues() ([]Queue, string) {
	var queues []Queue
	db, err := sqlx.Connect("mysql", dbURL)

	if err != nil {
		fmt.Println("Failed to connect", err)
		return queues, "Failed to connect"
	}
	defer db.Close()

	err = db.Select(&queues, "SELECT * FROM queue")
	if err != nil {
		fmt.Println("Failed to query all", err)
		return queues, "Failed to query all"
	}
	return queues, ""
}

func stringToBool(s string) bool {
	if s == "0" || s == "" {
		return false
	} else {
		return true
	}
}
