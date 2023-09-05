package ynab4

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"ynab4importer/internal/TB"
)

func WriteCsvFile(ofx TB.OFX, fileName string) {
	csvFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error creating CSV file:", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	// write header row
	fmt.Println("Date,Payee,Category,Memo,Outflow,Inflow")
	header := []string{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"}
	writer.Write(header)

	// write each STMTTRN row
	for _, trn := range ofx.Stmt.BankTransList.StmtTrns {
		if err := writeCsvRow(writer, trn); err != nil {
			log.Fatal("Error writing CSV row:", err)
		}
	}

	writer.Flush()
}

func writeCsvRow(writer *csv.Writer, trn TB.StmtTrn) error {
	memo := parseName(trn.Name)
	outflow, inflow := assignFlow(trn.TrnType, trn.TrnAmt)
	category := ParseCategory(memo)
	payee := ParsePayee(memo)
	date, err := formatDate(trn.DtAvail)
	if err != nil {
		return err
	}

	row := []string{date, payee, category, memo, outflow, inflow}
	if err := writer.Write(row); err != nil {
		return err
	}

	return nil
}

func parseName(input string) string {
	currencies := []string{"EUR ", "HUF ", "CZK "}
	for _, currency := range currencies {
		if idx := strings.Index(input, currency); idx >= 0 {
			return strings.TrimSpace(input[idx+len(currency):])
		}
	}
	return strings.TrimSpace(input)
}

func assignFlow(trnType string, trnAmt string) (outflow string, inflow string) {
	if trnType == "DEBIT" {
		outflow = trnAmt
	} else if trnType == "CREDIT" {
		inflow = trnAmt
	}
	return outflow, inflow
}

func formatDate(dateStr string) (string, error) {
	date, err := time.Parse("20060102", dateStr)
	if err != nil {
		return "", err
	}
	return date.Format("02/01/2006"), nil
}
