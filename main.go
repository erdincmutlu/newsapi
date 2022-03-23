package main

import (
	"flag"

	"github.com/erdincmutlu/newsapi/internal"
)

func main() {
	emailPass := flag.String("emailpass", "", "email password")
	dbPass := flag.String("dbpass", "", "db password")

	flag.Parse()

	internal.Start(*emailPass, *dbPass)
}
