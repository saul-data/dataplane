package main

import (
	"log"

	distributefilesystem "github.com/saul-data/dataplane/app/mainapp/code_editor/distribute_filesystem"
	dpconfig "github.com/saul-data/dataplane/app/mainapp/config"
	"github.com/saul-data/dataplane/app/mainapp/database"
)

func main() {
	dpconfig.LoadConfig()
	database.DBConnect()
	log.Println("🏃 Running")
	// CreateFiles()
	distributefilesystem.MoveCodeFilesToDB(database.DBConn)
}
