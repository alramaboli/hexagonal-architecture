package api

import "net/http"

//ProductHandler ...
type ProductHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
}
