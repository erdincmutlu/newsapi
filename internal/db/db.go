package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetFeed(dbClient *sqlx.DB, provider string, category string) (string, error) {
	var url string

	err := dbClient.Get(&url, "select url from feed where provider=? and category=?", provider, category)
	return url, err
}
