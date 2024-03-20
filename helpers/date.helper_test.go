package helpers

import "testing"

func TestConvertDate(t *testing.T) {
	date := "01-01-2020"
	result := ConvertDate(date)
	if result != "2020-01-01" {
		t.Error("Expected 2020-01-01, got ", result)
	}
}

func TestSplitFormNo(t *testing.T) {
	formNo := "001/2020"
	no, year, _ := SplitFormNo(formNo)
	if no != 1 || year != "2020" {
		t.Error("Expected 1, 2020, got ", no, year)
	}
}
