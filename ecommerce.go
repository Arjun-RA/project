package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const(
	username="root"
	password="Arjun@123"
	hostname="127.0.0.1:3306"
	dbName="project"
)



func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)

}

func update(data []string) sql.Result {
	fmt.Println("Connecting to db....")
	db, err := sql.Open("mysql", dsn("project"))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query, err_query := db.Prepare("UPDATE product_stock SET stock_qty =stock_qty + ? where product_id = ?;")
	if err_query != nil {
		fmt.Println(err)
	}
	defer query.Close()
	res, _ := query.Exec(data[1], data[0])
	return res

}

func main() {
	fmt.Println("This is a file reader")

	file, err := os.Open("Stock.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, val := range records {
		res := update(val)
		fmt.Println(res)
	}
}