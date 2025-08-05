package conf

import "net/http"

type RestError struct {
	Mensagen string `json:"mensagem"`
	Err      string `json:"err"`
	Code     int    `json:"code"`
	Causas   []Causa `json:causas`
}

type Causa struct {
	Field    string
	Mensagen string
}

func NewRestErr(mensagem, err string, code int, causas []Causa) *RestError {
	return &RestError{
		Mensagen: mensagem,
		Err:      err,
		Code:     code,
		Causas:   causas,
	}
}

func BadRequest_NewRestErr(mensagem string) *RestError {
	return &RestError{
		Mensagen: mensagem,
		Err:      "Bad request",
		Code:     http.StatusBadRequest,
	}
}

func BadRequest_ValidadeNil(mensagem string, causas []Causa) *RestError{
	return &RestError{
		Mensagen: mensagem,
		Err:      "Bad request",
		Code:     http.StatusBadRequest,
		Causas: causas,
	}
}

func InternoServeError_RestErr(mensagem string) *RestError{
	return &RestError{
		Mensagen: mensagem,
		Err:      "Interno serve error",
		Code:     http.StatusInternalServerError,
	}
}