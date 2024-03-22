package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

const port = 3306
const dbname = "sqldb"

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func initDB() error {
	var err error
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s", user, pass, port, dbname)
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	return nil
}

func createItemHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO items(name) VALUES(?)", item.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastInsertedID, _ := result.LastInsertId()
	item.ID = int(lastInsertedID)

	jsonItem, _ := json.Marshal(item)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonItem)
}

func getAllItemsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := db.Query("SELECT id, name FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	jsonItems, _ := json.Marshal(items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonItems)
}

func createSchemaFromDDLFile() error {
	file, err := os.ReadFile("Sample_table_ddl.sql")
	if err != nil {
		return err
	}

	ddlStatements := strings.Split(string(file), ";")
	for _, stmt := range ddlStatements {
		if strings.TrimSpace(stmt) == "" {
			continue
		}
		_, err := db.Exec(stmt)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateItemHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Implement update item handler
}

func deleteItemHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Implement delete item handler
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	err = createSchemaFromDDLFile()
	if err != nil {
		fmt.Println("Error creating schema from DDL file:", err)
		return
	}

	fmt.Println("Schema created successfully")

	router := httprouter.New()
	router.POST("/items", createItemHandler)
	router.GET("/items", getAllItemsHandler)
	router.PUT("/items/:id", updateItemHandler)
	router.DELETE("/items/:id", deleteItemHandler)

	fmt.Println("Running server on :8081")
	fmt.Println(http.ListenAndServe(":8081", router))
}
