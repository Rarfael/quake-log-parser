package main

import (
	"log"
	"os"

	"github.dom/Rarfael/quake-log-parser/parser"
	"github.dom/Rarfael/quake-log-parser/reports"
)

func main() {
	file, err := os.Open("logs/qgames.log")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()
	gameParser := parser.NewGameParser()
	games := gameParser.ParseLogFile(file)

	reports.GenerateReports(games)
}
