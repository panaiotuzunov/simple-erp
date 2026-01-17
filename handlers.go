package main

import (
	"net/http"
)

func (cfg *apiConfig) Test(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Add("Content-Type", "text/html; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("<html><body><h1>The server works!</h1></body></html>"))
}
