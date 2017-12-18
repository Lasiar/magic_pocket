package models

import (
	"database/sql"
	"magic_pocket/lib"
	"log"
)

func NewDb()  *sql.DB {
	var err error
	lib.Db, err = sql.Open("postgres", "postgres://vallder:30061997@192.168.0.102/mac_addr")
	if err != nil {
		log.Fatal(err)
	}
	err = lib.Db.Ping()
	if err != nil {
		log.Fatal()
	}
	return lib.Db
}
