package main

import (
	"fmt"
	"os"

	"github.com/Mirantis/statkube/db"
	"github.com/Mirantis/statkube/importer"
)

func main() {
	db := db.GetDB()
	filename, exists := os.LookupEnv("EMPLOYMENT_FILE")
	if !exists {
		filename = "default_data.json"
	}
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintln("Error opening file %s: %v ", filename, err.Error()))
	}
	importer.loadAll(db, f)
}
