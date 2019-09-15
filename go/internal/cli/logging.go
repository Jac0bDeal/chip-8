package cli

import (
	"log"
	"os"
)

func logAndExit(code int, msg string, a ...interface{}) {
	log.Printf(msg+"\n", a...)
	os.Exit(code)
}

func logErrorAndExit(err error) {
	logAndExit(1, err.Error())
}
