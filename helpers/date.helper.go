package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ConvertDate(date string) string {
	// Parse the input date string
	parsedDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return "2006-01-02"
	}

	return parsedDate.Format("2006-01-02")
}

func ConvertISOStringToDate(date string) string {
	// Parse the input date string
	parsedDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	return parsedDate.Format("2006-01-02")
}

func SplitFormNo(formNo string) (int, string, bool) {
	var no int
	var year string

	if !strings.Contains(formNo, "/") {
		no, _ = strconv.Atoi(formNo)
		year = formNo
		return no, year, false
	}

	fmt.Sscanf(formNo, "%d/%s", &no, &year)
	return no, year, true
}
