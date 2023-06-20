package main

import (
	"log"
	"os"
)

func main() {
	logFile, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	file, err := os.Open("no-existing-file.text")

	if err != nil {
		log.Println("error => ", err)
	}
	defer file.Close()
}
