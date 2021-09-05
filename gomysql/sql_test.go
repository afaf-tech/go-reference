package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES ('eko', 'EKO')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success inserting db")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	// loop data with rows.Next() return type bool
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}
	defer rows.Close()
	fmt.Println("")
}

func TestQueryComplexSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, rating, birth_date, balance, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	// loop data with rows.Next() return type bool
	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate, createdAt sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &rating, &birthDate, &balance, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("=========================")
		// cek if email is valid
		if email.Valid {
			fmt.Print("email valid	")
		} else {
			fmt.Print("email is null	")
		}
		fmt.Println(id, name, email, rating, birthDate, balance, married, createdAt)
	}
	defer rows.Close()
	fmt.Println("")
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// dont do this
	username := "admin'; #" // ; for end and # for comments. means will ignore the password
	password := "admin"

	// solution : dont create manual query with concat the string with user-input value.
	// will be described in the next chapter

	query := "SELECT username, password FROM user WHERE username= '" + username + "' AND password= '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	// loop data with rows.Next() return type bool
	if rows.Next() {
		var username, password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println(username, password)
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("gagal login")
	}
	defer rows.Close()
	fmt.Println("")
}
