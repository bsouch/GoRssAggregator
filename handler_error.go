package main

import "net/http"

func handlerError(writer http.ResponseWriter, request *http.Request) {
	errorJsonResponse(writer, 500, "Oops! Something went wrong. I won't tell you what it is because I'm a terrible API...")
}
