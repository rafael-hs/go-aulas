package main

import (
	"fmt"
	"main/logger" // importação direta
	// log "main/logger" <---- alias
	// . "main/logger" <---- dot import (não recomendado)
	// _ "main/logger" <---- init import
	"os"
)

func main() {

	logger.Log()

	logger.OtherLog()

	files, err := os.ReadDir(".")
	if err != nil {
		panic("error")
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
