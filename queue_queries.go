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

var dbHost string
var dbUsername string
var dbPassword string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	dbHost = os.Getenv("DB_HOST")
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
}

func insertQueue(q Queue) (Queue, string) {
	dbURL := getDBURL()
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

	dbURL := getDBURL()
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

func contactQueue(id int) (Queue, string) {
	var queue Queue

	dbURL := getDBURL()
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

	contacted := "1"
	if queue.Contacted {
		contacted = "0"
	}

	result := tx.MustExec("UPDATE queue SET contacted=? WHERE id=?", contacted, strconv.Itoa(id))
	rows, err := result.RowsAffected()

	if rows < 1 || err != nil {
		fmt.Println("Failed to update", err)
		return queue, "Failed to update"
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("Failed to commit transaction", err)
		return queue, "Failed to commit transaction"
	}

	queue.Contacted = !queue.Contacted
	return queue, ""
}

func getQueues() ([]Queue, string) {
	var queues []Queue
	dbURL := getDBURL()
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

func cleanDB() string {
	dbURL := getDBURL()
	db, err := sqlx.Connect("mysql", dbURL)
	if err != nil {
		fmt.Println("Failed to connect", err)
		return "Failed to connect"
	}
	defer db.Close()

	tx := db.MustBegin()
	tx.MustExec("TRUNCATE TABLE queue")

	return ""
}

func getDBURL() string {
	dbName := os.Getenv("DB_NAME")
	return dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":3306)" + "/" + dbName + "?parseTime=true"
}

func stringToBool(s string) bool {
	if s == "0" || s == "" {
		return false
	}
	return true
}
