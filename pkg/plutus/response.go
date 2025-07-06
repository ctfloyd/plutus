package plutus

import "github.com/ctfloyd/hazelmere-commons/pkg/hz_api"

type Meta struct {
	RequestId string `json:"request_id"`
	NextToken string `json:"next_token,omitempty"`
}

type ErrorResponse struct {
	hz_api.ErrorResponse
	Meta Meta `json:"meta"`
}

func (e ErrorResponse) GetStatus() int {
	return e.Status
}
