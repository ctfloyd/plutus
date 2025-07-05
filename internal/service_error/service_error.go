package service_error

import (
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_service_error"
	"net/http"
	"plutus/pkg/api"
)

var Internal = hz_service_error.ServiceError{Code: api.ErrorCodeInternal, Status: http.StatusInternalServerError}
var BadRequest = hz_service_error.ServiceError{Code: api.ErrorCodeBadRequest, Status: http.StatusBadRequest}
var Unauthorized = hz_service_error.ServiceError{Code: api.ErrorCodeUnauthorized, Status: http.StatusUnauthorized}
