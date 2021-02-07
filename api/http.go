package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/projects/mongodb/product"
)

//ProductHandler ...
type ProductHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
}

type handler struct {
	productService product.Service
}

//NewHandler ...
func NewHandler(productService product.Service) ProductHandler {

	return &handler{productService: productService}

}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := chi.URLParam(r, "code")
	p, err := h.productService.Find(code)

	if err != nil {

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	json.NewEncoder(w).Encode(&p)

}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	//requestBody, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	p := &product.Product{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = h.productService.Store(p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&p)

}
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	code := chi.URLParam(r, "code")
	err := h.productService.Delete(code)
	if err != nil {
		log.Fatal(err)
	}

}
func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	p, err := h.productService.FindAll()

	if err != nil {

		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&p)

}
