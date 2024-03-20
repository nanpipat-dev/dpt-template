package emsgs

import (
	"net/http"

	core "gitlab.finema.co/finema/idin-core"
)

var PermissionsDenied = core.Error{
	Status:  http.StatusForbidden,
	Code:    "PERMISSION_DENIED",
	Message: "permission denied!",
}

var AuthUsernameOrPasswordInvalid = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "INVALID_CREDENTIALS",
	Message: "username or password is an invalid",
}

var AuthUserInActive = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "PERMISSION_DENIED",
	Message: "this account is inactive",
}

var AuthCurrentPasswordInvalid = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "INVALID_CURRENT_PASSWORD",
	Message: "current password does not match",
}

var AuthTokenRequired = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "TOKEN_REQUIRED",
	Message: "Authorization header field is required"}

var VCQRTokenRequired = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "TOKEN_REQUIRED",
	Message: "X-Vc-Qr-Token header field is required"}

var AuthTokenInvalid = core.Error{
	Status:  http.StatusUnauthorized,
	Code:    "INVALID_TOKEN",
	Message: "Token is invalid"}

var AuthAPIKeyInvalid = core.Error{
	Status:  http.StatusUnauthorized,
	Code:    "INVALID_API_KEY",
	Message: "Api key is invalid"}

var AuthAPIKeyRequired = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "API_KEY_REQUIRED",
	Message: "x-api-key header field is required"}

var TransactionIDRequired = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "TRANSACTION_ID_REQUIRED",
	Message: "Transaction ID is required",
}

var TokenExpired = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "TOKEN_EXPIRED",
	Message: "Token is expired",
}

var DomainNotFound = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "DOMAIN_NOT_FOUND",
	Message: "User domain is not found",
}

var UserDeactivate = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "USER_DEACTIVATE",
	Message: "User status is deactivate",
}

var SamNotFound = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "ADFS_ACCOUNT_NOT_FOUND",
	Message: "adfs account not found",
}

var AuthPasswordAlreadySet = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "PASSWORD_ALREADY_SET",
	Message: "password already set",
}

var LineProfileNotFound = core.Error{
	Status:  http.StatusBadRequest,
	Code:    "LINE_PROFILE_NOT_FOUND",
	Message: "line profile not found. Please login again",
}
