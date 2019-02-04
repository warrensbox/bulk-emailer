package main_test

import (
	"os"
	"testing"
)

var test_csv string

func init() {
	test_csv = "test_content.txt"
}

func TestValidateOpenFileExist(t *testing.T) {

	csvFile, errorFile := os.Open(test_csv)
	defer csvFile.Close()

	if errorFile != nil {
		t.Errorf("Unable to open test case file : %s", test_csv)
	} else {
		t.Logf("Able to open file %s", test_csv)
	}
}
