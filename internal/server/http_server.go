package server

import (
	"Users/config"
	"net/http"
)

func InitHTTPServer(cfg *config.Config) *http.Server {
	return &http.Server{}
}
