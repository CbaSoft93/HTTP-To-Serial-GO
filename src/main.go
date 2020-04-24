package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./dto"
)

type server struct{}

func (s *server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		sendToSerial(res, req)
	default:
		resApi := dto.ResAPI{
			Status: http.StatusMethodNotAllowed,
			Result: "not allowed",
		}
		resApi.ToHTTPResponse(res)
	}
}

func sendToSerial(res http.ResponseWriter, req *http.Request) {
	var serialRequest dto.SerialRequest
	err := json.NewDecoder(req.Body).Decode(&serialRequest)
	var resApi dto.ResAPI

	if err != nil {
		resApi = dto.ResAPI{
			Status: http.StatusBadRequest,
			Result: err,
		}
	} else {
		resApi = serialRequest.SendToSerial()
	}

	resApi.ToHTTPResponse(res)
}

func main() {
	s := &server{}
	http.Handle("/", s)
	error := http.ListenAndServe(":6593", nil)
	log.Fatal(error)
}
