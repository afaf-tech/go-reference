package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "fikri"
	password := "fikri"
	query := "INSERT INTO customer(useranme, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
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

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// will be failed due to using param. not concatenation.
	username := "admin'; #" // ; for end and # for comments. means will ignore the password
	password := "admin"

	// solution : dont create manual query with concat the string with user-input value.
	// will be described in the next chapter

	query := "SELECT username, password FROM user WHERE username= ? AND password= ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)
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

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "paman'; DROP TABLES user"
	password := "paman"

	query := "INSERT INTO user(useranme, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success inserting db")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "fikri@gmail.com"
	comments := "nana lagi marah"

	query := "INSERT INTO comments(email, comments) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, email, comments)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success inserting db with id :", insertId)
}

// ExecContext and ExecQuery will renew the db connection after first execution. so that if
// we have more execution, the db connection will not be the same.
// But if we have statement, no matter how much execution we have. we just need one connection DB.
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comments(email, comments) VALUES (?, ?)"
	// prepare statement is commonly used when you have the same query and many params at the same time.
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "fikri" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment_" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("id :", id)
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction

	query := "INSERT INTO comments(email, comments) VALUES (?, ?)"
	for i := 0; i < 10; i++ {
		email := "fikri" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment_" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("id :", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
