package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"ynab4importer/internal/TB"
	"ynab4importer/internal/ynab4"
)

func main() {
	inputFileName := getInputFileName(os.Args)
	outputFileName := getOutputFileName(os.Args)

	log.Printf("Processing file - %s", inputFileName)
	log.Printf("Out file - %s", outputFileName)

	data, err := os.ReadFile(inputFileName)
	if err != nil {
		log.Fatalln("Error reading file:", err)
		return
	}

	var ofx TB.OFX
	err = xml.Unmarshal(data, &ofx)
	if err != nil {
		log.Fatalln("Error unmarshaling XML:", err)
		return
	}

	fmt.Printf("Writing %d transactions.\n", len(ofx.Stmt.BankTransList.StmtTrns))

	ynab4.WriteCsvFile(ofx, outputFileName)
}

func getInputFileName(args []string) string {
	if len(args) > 1 {
		return args[1]
	}
	return "download.ofx"
}

func getOutputFileName(args []string) string {
	if len(args) > 2 {
		return args[2]
	}
	return "output.csv"
}
