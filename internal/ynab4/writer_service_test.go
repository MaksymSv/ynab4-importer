package ynab4

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestFormatDate_ValidDateString(t *testing.T) {
	input := "20230512"
	expectedOutput := "12/05/2023"
	expectedError := error(nil)

	output, err := formatDate(input)

	if diff := cmp.Diff(expectedOutput, output); diff != "" {
		t.Errorf("formatDate(%s) output differs: (-want +got)\n%s", input, diff)
	}

	if err != expectedError {
		t.Errorf("formatDate(%s) error differs: expected %v, but got %v", input, expectedError, err)
	}
}

func TestFormatDate_InvalidDateFormat(t *testing.T) {
	input := "05/12/2023"
	expectedOutput := ""
	expectedError := &time.ParseError{}

	output, err := formatDate(input)

	if diff := cmp.Diff(expectedOutput, output); diff != "" {
		t.Errorf("formatDate(%s) output differs: (-want +got)\n%s", input, diff)
	}

	if reflect.TypeOf(err) != reflect.TypeOf(expectedError) {
		t.Errorf("formatDate(%s) error differs: expected %v, but got %v", input, expectedError, err)
	}
}

func TestFormatDate_InvalidDateValue(t *testing.T) {
	input := "20230000"
	expectedOutput := ""
	expectedError := &time.ParseError{}

	output, err := formatDate(input)

	if diff := cmp.Diff(expectedOutput, output); diff != "" {
		t.Errorf("formatDate(%s) output differs: (-want +got)\n%s", input, diff)
	}

	if reflect.TypeOf(err) != reflect.TypeOf(expectedError) {
		t.Errorf("formatDate(%s) error differs: expected %v, but got %v", input, expectedError, err)
	}
}

func TestAssignFlow_Debit(t *testing.T) {
	trnType := "DEBIT"
	trnAmt := "100.00"

	expectedOutflow := "100.00"
	expectedInflow := ""

	outflow, inflow := assignFlow(trnType, trnAmt)

	if outflow != expectedOutflow {
		t.Errorf("Unexpected outflow: expected=%q, got=%q", expectedOutflow, outflow)
	}

	if inflow != expectedInflow {
		t.Errorf("Unexpected inflow: expected=%q, got=%q", expectedInflow, inflow)
	}
}

func TestAssignFlow_Credit(t *testing.T) {
	trnType := "CREDIT"
	trnAmt := "200.00"

	expectedOutflow := ""
	expectedInflow := "200.00"

	outflow, inflow := assignFlow(trnType, trnAmt)

	if outflow != expectedOutflow {
		t.Errorf("Unexpected outflow: expected=%q, got=%q", expectedOutflow, outflow)
	}

	if inflow != expectedInflow {
		t.Errorf("Unexpected inflow: expected=%q, got=%q", expectedInflow, inflow)
	}
}

func TestAssignFlow_InvalidType(t *testing.T) {
	trnType := "INVALID"
	trnAmt := "300.00"

	expectedOutflow := ""
	expectedInflow := ""

	outflow, inflow := assignFlow(trnType, trnAmt)

	if outflow != expectedOutflow {
		t.Errorf("Unexpected outflow: expected=%q, got=%q", expectedOutflow, outflow)
	}

	if inflow != expectedInflow {
		t.Errorf("Unexpected inflow: expected=%q, got=%q", expectedInflow, inflow)
	}
}

func TestParseName_EUR(t *testing.T) {
	input := "20230426 14:00:25 21.17EUR Tesco  Bratislava Vien"

	expectedOutput := "Tesco  Bratislava Vien"

	output := parseName(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: expected=%q, got=%q", expectedOutput, output)
	}
}

func TestParseName_HUF(t *testing.T) {
	input := "tre 324HUF Jane Smith"

	expectedOutput := "Jane Smith"

	output := parseName(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: expected=%q, got=%q", expectedOutput, output)
	}
}

func BenchmarkParseName_HUF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Call the function being benchmarked here
		parseName("tre 324HUF Jane Smith")
	}
}

func TestParseName_CZK(t *testing.T) {
	input := "qwerty CZK Bob Johnson"

	expectedOutput := "Bob Johnson"

	output := parseName(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: expected=%q, got=%q", expectedOutput, output)
	}
}

func TestParseName_NoMatch(t *testing.T) {
	input := "USD Alice Brown"

	expectedOutput := "USD Alice Brown"

	output := parseName(input)

	if output != expectedOutput {
		t.Errorf("Unexpected output: expected=%q, got=%q", expectedOutput, output)
	}
}
