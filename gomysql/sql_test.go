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
