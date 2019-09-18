package main

import (
	"github.com/trangnnp-ts/assignment00/db"
	"github.com/trangnnp-ts/assignment00/router"
)

func main() {
	db.DB_connect()
	router.Routerdef()
}
