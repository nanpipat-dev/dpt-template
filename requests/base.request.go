package requests

import (
	"time"

	core "gitlab.finema.co/finema/idin-core"
)

func IsDate(input *string, fieldPath string) (bool, *core.IValidMessage) {
	if input == nil {
		return true, nil
	}

	DateFormat := "2006-01-02"
	_, ok := time.Parse(DateFormat, *input)
	if ok != nil {
		return false, DateM(fieldPath)
	}

	return true, nil
}

var DateM = func(field string) *core.IValidMessage {
	return &core.IValidMessage{
		Name:    field,
		Code:    "INVALID_DATE_TIME",
		Message: "The " + field + ` field must be in "yyyy-MM-dd" format`,
	}
}
