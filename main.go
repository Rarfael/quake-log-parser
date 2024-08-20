package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("logs/qgames.log")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()
	fmt.Println("File opened successfully")
}
