package main

import (
	"github.com/itsmontoya/aio"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
)

func main() {
	s := &srv{
		aio: aio.New(2),
	}

	r := httprouter.New()
	r.GET("/a", s.handleA)
	r.GET("/b", s.handleB)
	http.ListenAndServe(":1337", r)
}

type srv struct {
	aio *aio.AIO
}

func (s *srv) handleA(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	f, err := s.aio.Open("../../testing/declarationOfIndependence.txt")
	if err != nil {
		return
	}

	if _, err = io.Copy(w, f); err != nil {
		return
	}

	if err = f.Close(); err != nil {
		return
	}
}

func (s *srv) handleB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	f, err := os.Open("../../testing/declarationOfIndependence.txt")
	if err != nil {
		return
	}

	if _, err = io.Copy(w, f); err != nil {
		return
	}

	if err = f.Close(); err != nil {
		return
	}
}
