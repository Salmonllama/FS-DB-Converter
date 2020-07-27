package main

import (
	"database/sql"
	"fmt"
	"github.com/salmonllama/fs-db-converter/lib"
	"strings"
	"time"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqlDb, err := sql.Open("sqlite3", "./fsbot.db")
	if err != nil {
		fmt.Print(err)
	}
	createTable(sqlDb)

	rethinkDb, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "fsbot",
	})
	if err != nil {
		fmt.Print(err)
	}

	rows, err := r.Table("outfits").Run(rethinkDb)
	if err != nil {
		fmt.Print(err)
	}

	var outfits []lib.Outfit
	err = rows.All(&outfits)
	if err != nil {
		fmt.Print(err)
	}

	var count int = 0
	fmt.Printf("Starting with %v rows\n", len(outfits))
	for _, outfit := range outfits {
		if strings.Contains(outfit.Link, "i.imgur") {
			outfit.Created = time.Now().Unix()
			//fmt.Println(outfit.Link)
			outfit.Id = getIdFromLink(outfit.Link)

			statement, err := sqlDb.Prepare("INSERT INTO outfits (id, link, tag, submitter, created) VALUES (?, ?, ?, ?, ?)")
			if err != nil {
				fmt.Println(err)
			}

			_, err = statement.Exec(outfit.Id, outfit.Link, outfit.Tag, outfit.Submitter, outfit.Created)
			if err != nil {
				fmt.Println(err)
			}
			count += 1
		} else {
			// Do nothing, only want to keep imgur links
		}
	}
	fmt.Printf("Ended with %v rows\n", count)
}

func createTable(db *sql.DB) {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS outfits (id TEXT, link TEXT, submitter TEXT, tag TEXT, meta TEXT, created TEXT, updated TEXT, deleted TEXT, featured TEXT, display_count INT, delete_hash TEXT)")

	if err != nil {
		fmt.Print(err)
	}

	_, err = statement.Exec()
	if err != nil {
		fmt.Print(err)
	}
}

func getIdFromLink(link string) string {
	s := strings.Split(link, "/") // https:  i.imgur.com <id>.png
	id := strings.Split(s[3], ".")[0]
	return id
}
