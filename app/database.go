package app

import (
	"database/sql"
	"github.com/widadfjry/cashier-app/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:plmokn123@tcp(localhost:3306)/cashier_app")
	helper.PanicIfError(err)

	db.SetMaxOpenConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
