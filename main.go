package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wwwutz/elwms"
	"os"
	"path/filepath"
	"strconv"
)

func visit(path string, fi os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)

	return nil
}

func main() {

	database, _ := sql.Open("sqlite3", "omni-notes")
	// statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	// statement.Exec()

	flag.Parse()
	root := "."
	if flag.Arg(0) != "" {

		fi, err := os.Stat(flag.Arg(0))
//        fmt.Printf("err: %#v",err)
		elwms.Exiton(err, "supplied arg "+flag.Arg(0)+" failed")
        if !fi.IsDir() {
            fmt.Printf("// "+flag.Arg(0)+" is not a directory\n")
            os.Exit(1)
        }
		root = flag.Arg(0)
	}

	err := filepath.Walk(root, visit)

	fmt.Printf("filepath.Walk() returned %v\n", err)

	statement, _ := database.Prepare("INSERT INTO notes (creation, last_modification, title, content) VALUES (?, ?, ?, ?)")

	statement.Exec(861, 10000000, "Total Recall", "krasser film")
	rows, _ := database.Query("SELECT creation, title FROM notes")
	var mID int
	var title string
	for rows.Next() {
		rows.Scan(&mID, &title)
		fmt.Println(strconv.Itoa(mID) + ": " + title)
	}
}
