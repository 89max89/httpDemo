package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var database *sql.DB

type ViewData struct {
	TitleId      int
	TitleName    string
	TitleBalance int
}

func getId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	user := getUser(database, idd)
	d := ViewData{TitleId: user.id, TitleName: user.name, TitleBalance: user.balance}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, d)
}
