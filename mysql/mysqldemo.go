package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:dbstar@tcp(127.0.0.1:3306)/test3")
	if err != nil {
		log.Println(err)
		fmt.Println(err)
	}
	var str string
	name := "xuejun"
	rows, err := db.Query("SELECT name FROM class WHERE name=?", name)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var classname string
		if err := rows.Scan(&classname); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("get %s \n", classname)
	}
	err = db.QueryRow("SELECT * FROM class WHERE name=? ", name).Scan(&str)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
	}
	fmt.Println(str)
	defer db.Close()
	time.Sleep(30 * time.Second)
}
