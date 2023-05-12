package TB

import "encoding/xml"

type OFX struct {
	XMLName xml.Name `xml:"OFX"`
	Stmt    Stmt     `xml:"STMTRS"`
}

type Stmt struct {
	BankAcctFrom  BankAcctFrom  `xml:"BANKACCTFROM"`
	BankTransList BankTransList `xml:"BANKTRANLIST"`
}

type BankAcctFrom struct {
	BankID string `xml:"BANKID"`
	AcctID string `xml:"ACCTID"`
	IBAN   string `xml:"IBAN"`
}

type BankTransList struct {
	DtStart  string    `xml:"DTSTART"`
	DtEnd    string    `xml:"DTEND"`
	StmtTrns []StmtTrn `xml:"STMTTRN"`
}

type StmtTrn struct {
	TrnType      string     `xml:"TRNTYPE"`
	DtPosted     string     `xml:"DTPOSTED"`
	DtAvail      string     `xml:"DTAVAIL"`
	TrnAmt       string     `xml:"TRNAMT"`
	TrnVasym     string     `xml:"TRNVASYM"`
	ReferenceE2E string     `xml:"REFERENCE_E2E"`
	Name         string     `xml:"NAME"`
	BankAcctTo   BankAcctTo `xml:"BANKACCTTO"`
	Memo         string     `xml:"MEMO"`
	Currency     string     `xml:"CURRENCY"`
}

type BankAcctTo struct {
	BankID string `xml:"BANKID,omitempty"`
	AcctID string `xml:"ACCTID,omitempty"`
	IBAN   string `xml:"IBAN,omitempty"`
}
