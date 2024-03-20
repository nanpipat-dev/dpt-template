package emsgs

import (
	"net/http"

	core "gitlab.finema.co/finema/idin-core"
)

var EmailServiceSendFailed = core.Error{
	Status:  http.StatusBadGateway,
	Code:    "SEND_EMAIL_ERROR",
	Message: "can not send email",
}

var EmailServiceParserError = core.Error{
	Status:  http.StatusInternalServerError,
	Code:    "EMAIL_PARSING_ERROR",
	Message: "can not parse email",
}
