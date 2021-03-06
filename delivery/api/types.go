package api

import "github.com/ak-17/go-simple-web-api/usecase"

type Api struct {
	usecase usecase.Usecase
}

type Response struct {
	Header Header      `json:"header"`
	Data   interface{} `json:"data"`
}

type Header struct {
	ProcessTime  string   `json:"process_time"`
	ErrorMessage []string `json:"error_message"`
}
