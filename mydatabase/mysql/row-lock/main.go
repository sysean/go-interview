// 测试显示加行锁
// SELECT ... FOR UPDATE 之间会阻塞
// SELECT ... LOCK IN SHARE MODE 之间不会阻塞
// SELECT ... LOCK IN SHARE MODE 与 UPDATE 之间会阻塞

package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	SELECT_FOR_UPDATE_QUERY         = "SELECT * FROM students WHERE id = 1 FOR UPDATE"
	SELECT_LOCK_IN_SHARE_MODE_QUERY = "SELECT * FROM students WHERE id = 1 LOCK IN SHARE MODE"
	UPDATE_QUERY                    = "UPDATE students SET age = 10 WHERE id = 1"
)

type Student struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Age     int    `db:"age"`
	Gender  string `db:"gender"`
	Grade   string `db:"grade"`
	Address string `db:"address"`
}

func main() {
	defer fmt.Println("main() done")
	db, err := sqlx.Open("mysql", "sean:abc123@tcp(localhost:3306)/students_database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go anotherDbThread(&wg)

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	tx.Query(SELECT_FOR_UPDATE_QUERY)
	time.Sleep(10 * time.Second)
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		panic(err)
	}

	wg.Wait()
}

func anotherDbThread(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	defer fmt.Println("anotherDbThread() done")
	db, err := sqlx.Open("mysql", "sean:abc123@tcp(localhost:3306)/students_database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	tx.Query(SELECT_LOCK_IN_SHARE_MODE_QUERY)
	fmt.Println("anotherDbThread() do the query") // 看是否阻塞
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		panic(err)
	}
}
