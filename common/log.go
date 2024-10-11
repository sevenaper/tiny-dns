package common

import (
	"log"
	"os"
)

var file *os.File

func SetupLogger() error {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.SetOutput(file)
	return nil
}

func CloseLogger() error {
	return file.Close()
}
