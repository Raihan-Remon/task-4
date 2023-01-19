package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func DbConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:787898@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	return db
}
func postIndexHandler(w http.ResponseWriter, r *http.Request) {
	db := DbConnection()
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	user := User{Name: name, Email: email}
	fmt.Println(user)
	result, err := db.Exec("INSERT INTO user (name, email) VALUES (?, ?)", user.Name, user.Email)

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
	}
}

func getIndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := DbConnection()
	var users []User

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("users:", users)
	data, err := json.MarshalIndent(&users, "", "")
	w.Write(data)

}

func main() {
	db := DbConnection()
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS user (id INT AUTO_INCREMENT NOT NULL ,name TEXT NOT NULL, email TEXT NOT NULL, PRIMARY KEY (id))")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", getIndexHandler).Methods("GET")
	r.HandleFunc("/", postIndexHandler).Methods("POST")
	r.HandleFunc("/user", userHandler).Methods("GET")

	fmt.Println("Database Connected")
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8080", r))
}
