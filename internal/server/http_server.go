package server

import (
	"net/http"
)

func NewHTTPServer() *http.Server {
	return &http.Server{}
}
