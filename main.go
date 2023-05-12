package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"ynab4importer/internal/TB"
	"ynab4importer/internal/ynab4"
)

func main() {
	data, err := os.ReadFile("download.ofx")
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

	fmt.Printf("Bank ID: %s\n", ofx.Stmt.BankAcctFrom.BankID)
	fmt.Printf("Acc  ID: %s\n", ofx.Stmt.BankAcctFrom.AcctID)
	fmt.Printf("Transactions #: %d\n", len(ofx.Stmt.BankTransList.StmtTrns))

	writeCsvFile(ofx)
}

func writeCsvFile(ofx TB.OFX) {
	csvFile, err := os.Create("output.csv")
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

		memo := parseName(trn.Name)
		outflow, inflow := assignFlow(trn.TrnType, trn.TrnAmt)
		date, _ := formatDate(trn.DtAvail)
		category := ynab4.ParseCategory(memo)
		payee := ynab4.ParsePayee(memo)

		// line like: 31/01/2023,,,"DM Vienna Gate, BA",32.35,
		outLine := fmt.Sprintf(`%s,%s,"%s","%s",%s,%s`, date, payee, category, memo, outflow, inflow)
		fmt.Println(outLine)
		row := []string{date, payee, category, memo, outflow, inflow}
		writer.Write(row)
	}

	writer.Flush()
}

func parseName(input string) string {
	if parts := strings.Split(input, "EUR "); len(parts) > 1 {
		return parts[1]
	} else if parts := strings.Split(input, "HUF "); len(parts) > 1 {
		return parts[1]
	} else if parts := strings.Split(input, "CZK "); len(parts) > 1 {
		return parts[1]
	} else {
		return parts[0]
	}
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
