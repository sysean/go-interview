package main

import (
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// A 和 B 两个事务：
// A 事务先读取了 John 的信息，然后等待 B 事务结束，B 事务修改了 John 的信息，然后结束，A 事务再次读取 John 的信息。
// 在 READ COMMITTED 级别下，B 始终读
func TestRepeatedRead(t *testing.T) {
	db, err := sqlx.Open("mysql", "sean:abc123@tcp(localhost:3306)/students_database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		// 在这里进行事务 A 的操作
		tx, err := db.Beginx()
		if err != nil {
			panic(err)
		}

		ch2 <- struct{}{} // 通知事务 B 开始

		<-ch1 // 等待事务 B 读取

		sql1 := "UPDATE students SET grade = 'C' WHERE name = 'John'"
		_, err = tx.Exec(sql1)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		fmt.Println("Transaction A write:", sql1)
		fmt.Println("")

		if err := tx.Commit(); err != nil {
			tx.Rollback()
			panic(err)
		}
		fmt.Println("Transaction A commit")

		ch2 <- struct{}{} // 通知事务 B 结束
	}()

	<-ch2

	tx, err := db.Beginx()
	if err != nil {
		panic(err)
	}

	read := func() {
		var st []*Student
		sql1 := "SELECT * FROM students WHERE name = 'John'"
		if err := tx.Select(&st, sql1); err != nil {
			tx.Rollback()
			panic(err)
		}

		fmt.Println("Transaction B read:", sql1)
		fmt.Printf("st: %+v\n", *st[0])
		fmt.Println("")
	}

	read()

	ch1 <- struct{}{} // 通知事务 A 开始

	<-ch2 // 等待事务 A 结束

	read()

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		panic(err)
	}
	fmt.Println("Transaction B commit")

	fmt.Println("done")
}
