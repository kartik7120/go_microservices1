package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Goodbye World")
}
