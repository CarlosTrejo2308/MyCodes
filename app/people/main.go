package main

import (
	"github.com/CarlosTrejo2308/peopleApiResource/db"
)

func main() {
	// The name of the file where the people data is stored
	path := "people.xml"
	people := readFile(path)

	// Get the DB path and connect to it
	uripath := db.GeneratePath()
	db := db.Connect(uripath)

	insertToBd(people, db)
}
