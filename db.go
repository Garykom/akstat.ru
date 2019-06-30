package main

import (
	//"database/sql"
	//"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	//"math/rand"
	//"net/http"
	//b "rest/entities"
	//"bufio"
	//"encoding/binary"
	//"net"
	//"os"
	//"strconv"
	//utils "1cvpn_conf/utils"
	//"math/big"
	//"sort"
	//"strings"
	"time"
)

var (
	pool         *sqlx.DB
	dbErr        error
	dbConnString = "admin:{frfn0y@tcp(127.0.0.1:3306)/akstat?charset=utf8"
)

func init() {
	//pool, err = sqlx.Open("postgres", "postgres://root:root@postgresql-server.dev/wiki?sslmode=disable")
	pool, dbErr = sqlx.Open("mysql", dbConnString)
	pool.SetMaxIdleConns(5)
	pool.SetConnMaxLifetime(2 * time.Minute)
	pool.SetMaxOpenConns(95)
	if dbErr != nil {
		log.Println("m=GetPool,msg=connection has failed", dbErr)
		fmt.Println(dbErr)
	}
}

func GetConnection() (*sqlx.DB, error) {
	return pool, dbErr
}
