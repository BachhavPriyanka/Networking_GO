package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(149.28.24.60:6603)/ecommerceWebsite")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/products", productsHandler)

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	json.NewEncoder(w).Encode(products)
}
