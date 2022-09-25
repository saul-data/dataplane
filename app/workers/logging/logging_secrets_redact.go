package logging

import (
	"fmt"
	"log"

	wrkerconfig "github.com/saul-data/dataplane/app/workers/config"
)

func PrintSecretsRedact(logmessages ...interface{}) {

	// 	Secrets.Replace(logmessages...)
	var printString string
	for _, n := range logmessages {
		printString = printString + fmt.Sprintf("%s ", n)

	}

	log.Println(wrkerconfig.Secrets.Replace(printString))
}
