package main

import (
	"database/sql"
	"fmt"
	"github.com/salmonllama/fs-db-converter/lib"
	"strings"

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
	rows.All(&outfits)
	fmt.Println(getIdFromLink(outfits[0].Link))
	//for _, outfit := range outfits {
	//
	//}
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
