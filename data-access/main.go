// go mod init data-access
// create reverse.go
// add line in go.work ./moduleNEW
// go run ./data-access/
// https://edwinsiby.medium.com/performing-crud-operations-in-postgresql-with-go-42657761125c

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	var choice int
	db := connectPostgresDB()
	for {
		fmt.Println("Choose:\n 1 - Insert\n 2 - Read\n 3 - Update\n 4 - Delete")
		_, err := fmt.Scan(&choice)
		if err != nil {
			return
		}

		switch choice {
		case 1:

			Insert(db)
		case 2:
			Read(db)
		case 3:
			Update(db)
		case 4:
			Delete(db)
		}

	}

}

func Delete(db *sql.DB) {
	var id int
	fmt.Println("Enter ID to DELETE: ")
	_, err1 := fmt.Scan(&id)
	if err1 != nil {
		return
	}

	_, err := db.Query("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User deleted")
	}
}

func Update(db *sql.DB) {
	var id int
	var new_name string
	fmt.Println("Enter ID: ")
	_, err1 := fmt.Scan(&id)
	if err1 != nil {
		return
	}
	fmt.Println("Enter new name: ")
	_, err2 := fmt.Scan(&new_name)
	if err2 != nil {
		return
	}
	_, err := db.Query("UPDATE users SET name=$1 WHERE id=$2", new_name, id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated successfully!")
	}

}

func Read(db *sql.DB) {
	ReadFromPostgress(db)
}

func ReadFromPostgress(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	fmt.Println("id  |  name  |  email")
	fmt.Println("---------------------")

	var id int
	var name string
	var email string

	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%d  |  %s  |  %s\n", id, name, email)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
}

func Insert(db *sql.DB) {
	name := "Bob"
	email := "bob@m.re"
	InsertIntoPostgres(db, name, email)
}

func InsertIntoPostgres(db *sql.DB, name string, email string) {
	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	if err != nil {
		fmt.Println("Insert error:", err)
	} else {
		fmt.Println("Successfully inserted")
	}
}

func connectPostgresDB() *sql.DB {
	connstring := "user=postgres password=admin dbname=test host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
