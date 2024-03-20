package emsgs

import (
	"net/http"

	core "gitlab.finema.co/finema/idin-core"
)

var RequesterError = core.Error{
	Status:  http.StatusForbidden,
	Code:    "REQUESTER_ERROR",
	Message: "requester error!",
}

var AssociationNameEnIsAlreadyExist = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "ASSOCIATION_NAME_EN_IS_ALREADY_EXIST",
	Message: "association name en is already exist!",
}

var AssociationNameThIsAlreadyExist = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "ASSOCIATION_NAME_TH_IS_ALREADY_EXIST",
	Message: "association name th is already exist!",
}
