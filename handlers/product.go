package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kartik7120/go_microservices1/data"
)

type Product struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

/*
 Don't Need this code anymore
*/

// func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		p.l.Print("PUT request")
// 		reg := regexp.MustCompile("/([0-9]+)")

// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(g) != 1 {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		if len(g[0]) != 2 {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		idString := g[0][1]
// 		id, err := strconv.Atoi(idString)

// 		if err != nil {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		p.updateProducts(id, rw, r)

// 	}

// 	// Catch all

// 	rw.WriteHeader(http.StatusMethodNotAllowed)
// }

func (p *Product) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	// handle the request
	lp := data.GetProducts()
	// d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Product) AddProducts(rw http.ResponseWriter, r *http.Request) {
	// p.l.Println("Handle POST request")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	// What does this do?
	// p.l.Printf("Prod: %#v", prod) // %#v is a Go syntax to print the struct

	// p.l.Printf("%#v", prod)

	data.AddProducts(&prod)
}

func (p *Product) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT request")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	e := data.UpdateProducts(id, &prod)

	if e == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if e != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

type KeyProduct struct{}

func (p Product) MiddleWareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
