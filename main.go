package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wwwutz/elwms"
	"os"
	"strconv"
)

func visit(path string, fi os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

func readDirNames(dirname string) ([]string, error) {
	d, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	return names, nil
}

func main() {
	database, err := sql.Open("sqlite3", "omni-notes")
	elwms.Exiton(err, "sql.Open omni-notes failed")
	defer database.Close()
	// statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	// statement.Exec()
	//        fmt.Printf("err: %#v",err)

	flag.Parse()
	root := "."
	if flag.Arg(0) != "" {

		fi, err := os.Stat(flag.Arg(0))
		elwms.Exiton(err, "supplied arg "+flag.Arg(0)+" failed")
		if !fi.IsDir() {
			fmt.Printf("// " + flag.Arg(0) + " is not a directory\n")
			os.Exit(1)
		}
		root = flag.Arg(0)
	}
	// readDirNames reads the directory named by dirname and returns

	tree, err := readDirNames(root)
	elwms.Exiton(err, "readDirNames("+root+") failed")
	for i, entry := range tree {
		fmt.Printf("-%d-> %s\n", i, entry)
	}

	statement, err := database.Prepare("INSERT INTO notes (creation, last_modification, title, content) VALUES (?, ?, ?, ?)")
	elwms.Exiton(err, "db.Prepare() failed")

	statement.Exec(861, 10000000, "Total Recall", "krasser film")
	rows, err := database.Query("SELECT creation, title FROM notes")
	elwms.Exiton(err, "db.Query() failed")
	var mID int
	var title string
	for rows.Next() {
		rows.Scan(&mID, &title)
		fmt.Println(strconv.Itoa(mID) + ": " + title)
	}
}
