package service_error

import (
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_service_error"
	"net/http"
	"plutus/pkg/plutus"
)

var Internal = hz_service_error.ServiceError{Code: plutus.ErrorCodeInternal, Status: http.StatusInternalServerError}
var BadRequest = hz_service_error.ServiceError{Code: plutus.ErrorCodeBadRequest, Status: http.StatusBadRequest}
var Unauthorized = hz_service_error.ServiceError{Code: plutus.ErrorCodeUnauthorized, Status: http.StatusUnauthorized}
