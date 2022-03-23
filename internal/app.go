package internal

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var cache ttlcache.SimpleCache
var emailPass string
var dbClient *sqlx.DB
var dbPasswd string

const (
	port = "8080"

	// DB Settings
	dbUser = "root"
	dbAddr = "localhost"
	dbPort = "3306"
	dbName = "news"
)

func Start(ePassword, dbPassword string) {
	// Set up cache
	cache = ttlcache.NewCache()
	cache.SetTTL(time.Duration(10 * time.Second))

	// Parameters
	emailPass = ePassword
	dbPasswd = dbPassword

	// DB Client
	dbClient = getDbClient()

	// Set up mux router
	router := mux.NewRouter()
	router.
		HandleFunc("/news", getNews).
		Methods(http.MethodGet).
		Name("GetNews")

	router.
		HandleFunc("/share", shareNews).
		Methods(http.MethodPost).
		Name("ShareNews")

	log.Printf("starting service on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getDbClient() *sqlx.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
