package main

import "testing"

func TestGetInputFileName_HasTwoParams(t *testing.T) {
	// Test with arguments passed
	args := []string{"./program", "input.ofx", "output.csv"}
	inputFileName := getInputFileName(args)
	if inputFileName != "input.ofx" {
		t.Errorf("Expected input file name 'input.ofx', but got '%s'", inputFileName)
	}
}

func TestGetInputFileName_HasOneParam(t *testing.T) {
	// Test with arguments passed
	args := []string{"./program", "input.ofx"}
	inputFileName := getInputFileName(args)
	if inputFileName != "input.ofx" {
		t.Errorf("Expected input file name 'input.ofx', but got '%s'", inputFileName)
	}
}

func TestGetInputFileName_HasZeroParam(t *testing.T) {
	// Test with arguments passed
	args := []string{"./program"}
	inputFileName := getInputFileName(args)
	if inputFileName != "download.ofx" {
		t.Errorf("Expected defailt file name 'download.ofx', but got '%s'", inputFileName)
	}
}

func TestGetOutputFileName_HasTwoParams(t *testing.T) {
	// Test with arguments passed
	args := []string{"./program", "input.ofx", "my_output.csv"}
	fileName := getOutputFileName(args)
	if fileName != "my_output.csv" {
		t.Errorf("Expected output file name 'my_output.csv', but got '%s'", fileName)
	}
}

func TestGetOutputFileName_HasOneParams(t *testing.T) {
	// Test with arguments passed
	args := []string{"./program", "input.ofx"}
	fileName := getOutputFileName(args)
	if fileName != "output.csv" {
		t.Errorf("Expected defailt file name 'output.csv', but got '%s'", fileName)
	}
}

func TestGetOutputFileName_HasZeroParams(t *testing.T) {
	// Test with arguments passed
	args := []string{"./program"}
	fileName := getOutputFileName(args)
	if fileName != "output.csv" {
		t.Errorf("Expected defailt file name 'output.csv', but got '%s'", fileName)
	}
}
