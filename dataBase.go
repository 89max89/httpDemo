package main

import (
	"database/sql"
	"fmt"
	"log"
)

type user struct {
	id      int
	name    string
	balance int
}

func getUser(db *sql.DB, id int) *user {
	row := db.QueryRow("select * from  buyers where id = $1", id)
	user := user{}
	err := row.Scan(&user.id, &user.name, &user.balance)
	if err != nil {
		log.Fatalf("getting user error")
	}
	fmt.Println("buyer : ", user.name, "  balance : ", user.balance)
	return &user
}
