package dto

import (
	"encoding/json"
	"net/http"
)

// ResAPI is a DTO
type ResAPI struct {
	Status int
	Result interface{}
}

// ToHTTPResponse write ResAPI in the response
func (rAPI *ResAPI) ToHTTPResponse(res http.ResponseWriter) {
	json, err := json.Marshal(rAPI)
	if json != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(rAPI.Status)
		res.Write(json)
	} else {
		resAPIErr := ResAPI{
			Status: http.StatusInternalServerError,
			Result: err,
		}
		resAPIErr.ToHTTPResponse(res)
	}
}
