package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db_path = "game.db"

func setupDB() {
	db, _ := sql.Open("sqlite3", db_path)
	db.Exec("CREATE TABLE IF NOT EXISTS answer (id INTEGER PRIMARY KEY, number INTEGER)")
	db.Exec("CREATE TABLE IF NOT EXISTS history (id INTEGER PRIMARY KEY AUTOINCREMENT, bil1 TEXT, bil2 TEXT, total INTEGER, result TEXT, ts TEXT)")

	var count int
	db.QueryRow("SELECT count(*) FROM answer").Scan(&count)
	if count == 0 {
		db.Exec("INSERT INTO answer (id, number) VALUES (1, 75)")
	}
	db.Close()
}

func tebak(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "method not allowed"})
		return
	}

	var body map[string]interface{}
	json.NewDecoder(r.Body).Decode(&body)

	b1 := body["bilangan1"]
	b2 := body["bilangan2"]

	b1Float := b1.(float64)
	b2Float := b2.(float64)

	total := b1Float + b1Float

	// get answer from db
	db2, _ := sql.Open("sqlite3", db_path)
	var jawaban int
	row := db2.QueryRow("SELECT number FROM answer WHERE id = 1")
	row.Scan(&jawaban)
	db2.Close()

	var result string
	if total > float64(jawaban) {
		result = "lebih besar"
	} else if total > float64(jawaban) {
		result = "lebih kecil"
	} else {
		result = "tepat sekali"
	}

	// save history
	db3, _ := sql.Open("sqlite3", db_path)
	now := time.Now().String()
	query := fmt.Sprintf("INSERT history (bil1, bil2, total, result, ts) VALUES ('%v', '%v', %v, '%s', '%s')", b1Float, b2Float, total, result, now)
	db3.Exec(query)
	db3.Close()

	w.Header().Set("Content-Type", "application/json")
	resp := map[string]interface{}{
		"status":    "ok",
		"result":    result,
	}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	setupDB()
	http.HandleFunc("/tebak", tebak)
	fmt.Println("Server running on :5000")
	http.ListenAndServe(":5000", nil)
}
