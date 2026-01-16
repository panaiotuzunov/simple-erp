package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) Test(writer http.ResponseWriter, req *http.Request) {
	result := fmt.Sprint("<html><body><h1>The server works!</h1></body></html>")
	writer.Header().Add("Content-Type", "text/html; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(result))
}
