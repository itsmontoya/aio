package main

import (
	"io"
	"net/http"
	"os"

	"github.com/itsmontoya/async"
	"github.com/julienschmidt/httprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	s := &srv{
		a: async.New(2),
	}

	go func() {
		r := httprouter.New()
		r.GET("/a", s.handleA)
		r.GET("/b", s.handleB)
		http.ListenAndServe(":1337", r)
	}()

	go func() {
		fasthttp.ListenAndServe(":8081", s.handleC)
	}()

	select {}
}

type srv struct {
	a *async.Async
}

func (s *srv) handleA(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	f, err := s.a.Open("../../testing/declarationOfIndependence.txt")
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

func (s *srv) handleC(ctx *fasthttp.RequestCtx) {
	f, err := s.a.Open("../../testing/declarationOfIndependence.txt")
	if err != nil {
		return
	}

	if _, err = io.Copy(ctx, f); err != nil {
		return
	}

	if err = f.Close(); err != nil {
		return
	}
}
