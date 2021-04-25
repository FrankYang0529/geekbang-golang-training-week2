package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/FrankYang0529/geekbang-golang-training-week2/storage"
)

func main() {
	// init db
	mysqlURL := os.Getenv("mysql")
	db, err := sql.Open("mysql", mysqlURL)
	if err != nil {
		log.Fatalf("main: can't connect mysql, err: %+v", err)
	}
	defer db.Close()

	// init item storage
	itemStorage, err := storage.NewItemStorage(db)
	if err != nil {
		log.Fatalf("main: can't init item storage, err: %+v", err)
	}

	// get item
	item, err := itemStorage.GetItem("testItemID")
	if err != nil {
		log.Fatalf("main: can't get item, err: %+v", err)
	}
	log.Printf("item: %+v", item)
}
