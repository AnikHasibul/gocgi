package gocgi

import (
	"net/http"
	"net/http/cgi"
)

// CGIfy serves http with php script(cgi mode)
func CGIfy(w http.ResponseWriter, r *http.Request, cgiBin string, scriptFile string) {
	handler := new(cgi.Handler)
	handler.Path = cgiBin
	handler.Env = append(handler.Env, "REDIRECT_STATUS=CGI")
	handler.Env = append(handler.Env, "SCRIPT_FILENAME="+scriptFile)

	handler.ServeHTTP(w, r)
}
