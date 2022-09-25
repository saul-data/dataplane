package main

import (
	"log"

	wrkerconfig "github.com/saul-data/dataplane/app/workers/config"
	"github.com/saul-data/dataplane/app/workers/database"
)

func main() {
	wrkerconfig.LoadConfig()
	database.DBConnect()
	log.Println("🏃 Running")
	// CreateFiles()
	// distfilesystem.DownloadFiles()
}
