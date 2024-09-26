package server

import (
	"net/http"

	"Users/config"
)

func InitHTTPServer(cfg *config.Config) *http.Server {
	return &http.Server{}
}
