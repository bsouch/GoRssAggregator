package main

import "net/http"

func handlerReadiness(writer http.ResponseWriter, request *http.Request) {
	jsonResponse(writer, 200, struct{}{})
}
