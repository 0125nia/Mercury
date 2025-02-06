package ipconf

import "github.com/0125nia/Mercury/ipconf/domain"

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

// Response is the response struct
func ipConfResp(ed []*domain.Endpoint) Response {
	return Response{
		Message: "ok",
		Code:    0,
		Data:    ed,
	}
}

// top5Endpoints returns the top 5 endpoints
func top5Endpoints(eds []*domain.Endpoint) []*domain.Endpoint {
	if len(eds) <= 5 {
		return eds
	}
	return eds[:5]
}
