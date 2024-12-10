package handlers

import "net/http"

func HandleError(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}
