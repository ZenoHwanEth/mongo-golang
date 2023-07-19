package main

import (
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.GET("")
	r.POST("")
	r.DELETE("")
}
