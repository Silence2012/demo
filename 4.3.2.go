/*
test
 */
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"time"
)

var MAX_POOL_SIZE = 20
var DbPool chan *sql.DB

const (
	user     = "root"
	password = "secret"
	dbhost   = "127.0.0.1:3306"
	dbName   = "scmdb"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func CreatePool(db *sql.DB) {
	if DbPool == nil {
		DbPool = make(chan *sql.DB, MAX_POOL_SIZE)
	}
	if len(DbPool) > 20 {
		db.Close()
		return
	}
	DbPool <- db
}

func InitDb() {
	if DbPool == nil || len(DbPool) == 0 {
		DbPool = make(chan *sql.DB, MAX_POOL_SIZE)
		go func() {
			for i := 0; i < MAX_POOL_SIZE; i++ {
				dataSource := user + ":" + password + "@tcp(" + dbhost + ")/" + dbName + "?charset=utf8"
				//log.Println(dataSource)
				db, err := sql.Open("mysql", dataSource)
				if err != nil {
					log.Panicln(err)
				}
				CreatePool(db)
			}
		}()
	}
}
func GetDb() *sql.DB {
	if DbPool == nil || len(DbPool) == 0 {
		InitDb()
	}
	return <-DbPool
}
// test main
func main() {

	//log.Println("haha")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 5; i++ {
		OpDb(r.Intn(50))
		time.Sleep(time.Second)
		//log.Trace("good")
		//log.Trace(r.Intn(50))
	}
}

func OpDb(num int) {

	db := GetDb()
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	defer tx.Commit()
	rows, err := tx.Query("select ip,eidc  from ark limit 1")
	defer rows.Close()
	for rows.Next() {
		var num,num2 string
		if err := rows.Scan(&num,&num2); err != nil {
			log.Println(err)
		}
		log.Println(num,num2)
	}
	if err != nil {
		log.Println(err)
	}
}
