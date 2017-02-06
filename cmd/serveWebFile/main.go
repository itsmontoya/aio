package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/OneOfOne/aio"
	"github.com/julienschmidt/httprouter"
)

func main() {
	s := &srv{
		aio: aio.NewNum(2),
		// aio: aio.New(), using numCpu slows it down
	}

	r := httprouter.New()
	r.GET("/a", s.handleA)
	r.GET("/b", s.handleB)
	log.Fatal(http.ListenAndServe(":1337", r))
}

type srv struct {
	aio *aio.AIO
}

func (s *srv) handleA(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	or := <-s.aio.Open("../../testing/declarationOfIndependence.txt")
	if or.Err != nil {
		return
	}

	if _, err := io.Copy(w, or.F); err != nil {
		return
	}

	if err := or.F.Close(); err != nil {
		return
	}
}

func (s *srv) handleB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	f, err := os.Open("../../testing/declarationOfIndependence.txt")
	if err != nil {
		return
	}

	if _, err := io.Copy(w, f); err != nil {
		return
	}

	if err := f.Close(); err != nil {
		return
	}
}
