package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const DatabaseURL = "root:ironman.abc@tcp(localhost:3306)/tysql?charset=utf8"

type Customers struct {
	id   string
	name string
}

func main() {
	db, err := sql.Open("mysql", DatabaseURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("open db ok")

	rows, err := db.Query("select cust_id, cust_name from Customers")
	if err != nil {
		panic(err)
	}

	var customers []Customers
	for rows.Next() {
		// var cust_id string
		// var cust_name string
		var cus Customers

		if err := rows.Scan(&cus.id, &cus.name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(cus.id, cus.name)
		customers = append(customers, cus)
	}

	fmt.Println(customers)

	stmt, err := db.Prepare("insert Customers set cust_id=?, cust_name=?")
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec("1000000006", "kevin")

	db.Close()
}
