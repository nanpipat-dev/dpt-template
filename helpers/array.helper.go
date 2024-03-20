package helpers

import (
	"bytes"
	"reflect"
	"strings"

	csvs "encoding/csv"
	"encoding/json"

	zerowidth "github.com/trubitsyn/go-zero-width"
	core "gitlab.finema.co/finema/idin-core"
)

func ArrayMap[T any, R any](raws []T, cb func(r T) R) []R {
	items := make([]R, 0)

	for _, r := range raws {
		items = append(items, cb(r))
	}

	return items
}

func ReadFromFileMaps(data []byte, options *core.ICSVOptions) ([]map[string]interface{}, error) {
	reader := bytes.NewReader(data)
	csvRead := csvs.NewReader(reader)
	csvRead.LazyQuotes = true
	csvData, err := csvRead.ReadAll()
	if err != nil {
		return nil, err
	}

	dataList := []map[string]interface{}{}
	var dataLineOne []string

	for i, line := range csvData {
		if i == 0 {
			dataLineOne = line
		} else {
			data := map[string]interface{}{}
			for i2, line2 := range dataLineOne {
				clean := zerowidth.RemoveZeroWidthCharacters(line2)
				data[strings.ReplaceAll(clean, "\"", "")] = line[i2]
			}
			dataList = append(dataList, data)
		}
	}

	return dataList, nil
}

func CompareJSON(data1, data2 string) bool {
	var obj1 map[string]interface{}
	var obj2 map[string]interface{}

	// Unmarshal JSON strings into objects
	json.Unmarshal([]byte(data1), &obj1)
	json.Unmarshal([]byte(data2), &obj2)

	// Deep equal comparison
	if reflect.DeepEqual(obj1, obj2) {
		return true
	} else {
		// Compare attachments codes
		attachments1 := obj1["attachments"].([]interface{})
		attachments2 := obj2["attachments"].([]interface{})
		// Create maps to track attachment codes in obj2
		attachmentCodes := make(map[string]bool)
		for _, attachment := range attachments2 {
			code := attachment.(map[string]interface{})["code"].(string)
			attachmentCodes[code] = true
		}

		// Check if the codes for "additional" attachments differ in obj1 and obj2
		for _, attachment := range attachments1 {
			code := attachment.(map[string]interface{})["code"].(string)
			if code == "additional" {
				if _, ok := attachmentCodes[code]; !ok {
					return false
				}
			}
		}

		return true
	}

}

// // CompareData compares two sets of data
// func CompareData(data1, data2 string) bool {
// 	// Convert data structures to JSON for comparison
// 	data1JSON, _ := json.Marshal(data1)
// 	data2JSON, _ := json.Marshal(data2)

// 	// Compare JSON strings
// 	if string(data1JSON) != string(data2JSON) {
// 		return false
// 	}

// 	// Compare attachments
// 	for i := 0; i < len(data1.Attachments); i++ {
// 		if data1.Attachments[i].Code == "additional" {
// 			continue
// 		}

// 		// Find corresponding attachment in data2
// 		found := false
// 		for j := 0; j < len(data2.Attachments); j++ {
// 			if reflect.DeepEqual(data1.Attachments[i], data2.Attachments[j]) {
// 				found = true
// 				break
// 			}
// 		}

// 		// If corresponding attachment not found in data2, return false
// 		if !found {
// 			return false
// 		}
// 	}

// 	return true
// }

func ConvertStructToJsonString(data interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}
